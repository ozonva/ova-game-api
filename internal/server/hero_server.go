package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-game-api/internal/kafka"
	"github.com/ozonva/ova-game-api/internal/metrics"
	"github.com/ozonva/ova-game-api/internal/repo"
	"github.com/ozonva/ova-game-api/internal/saver"
	"github.com/ozonva/ova-game-api/internal/utils"
	"github.com/ozonva/ova-game-api/pkg/game"
	api "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const chunkSize = 2

type HeroServer struct {
	api.UnimplementedHeroApiServer
	logger        *zerolog.Logger
	repository    repo.HeroRepo
	saver         saver.Saver
	metrics       metrics.Metrics
	kafkaProducer kafka.KafkaProducer
}

func NewHeroApiServer(logger *zerolog.Logger, repo repo.HeroRepo, saver saver.Saver, metrics metrics.Metrics, kafkaProducer kafka.KafkaProducer) api.HeroApiServer {
	return &HeroServer{
		UnimplementedHeroApiServer: api.UnimplementedHeroApiServer{},
		logger:                     logger,
		repository:                 repo,
		saver:                      saver,
		metrics:                    metrics,
		kafkaProducer:              kafkaProducer,
	}
}

func (s *HeroServer) MultiCreateHero(ctx context.Context, request *api.MultiCreateHeroRequest) (*emptypb.Empty, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateHero")
	defer parentSpan.Finish()

	s.logger.Info().Msgf("MultiCreateHero request: %v", request)
	listHeroes := make([]game.Hero, 0, len(request.Heroes))

	for _, reqHero := range request.Heroes {
		typeHero := game.SearchTypeHeroesEnums(reqHero.TypeHero)
		hero := game.NewHero(reqHero.UserId, typeHero, reqHero.Name)
		hero.SetDescription(reqHero.Description)
		listHeroes = append(listHeroes, hero)
	}

	list, err := utils.HeroesToChunks(listHeroes, chunkSize)
	if err != nil {
		s.logger.Info().Msgf("MultiCreateHero request HeroesToChunks: %s", err)
		return nil, err
	}
	for _, chunkHeroes := range list {
		if err := s.saveHeroesChunks(ctx, parentSpan, chunkHeroes); err != nil {
			return nil, err
		}
		err := s.kafkaProducer.Send(kafka.Message{
			MessageType: kafka.MultiCreate,
			Value:       chunkHeroes,
		})
		if err != nil {
			s.logger.Info().Msgf("Can not send event to kafka, error: %s", err)
			return nil, err
		}
		s.metrics.AddSuccessMultiCreateHeroesCounter(float64(len(chunkHeroes)))
	}

	return &emptypb.Empty{}, nil
}

func (s *HeroServer) CreateHero(ctx context.Context, request *api.CreateHeroRequest) (*api.CreateHeroResponse, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "CreateHero")
	defer parentSpan.Finish()
	s.logger.Info().Msgf("CreateHero request: %v", request)

	typeHero := game.SearchTypeHeroesEnums(request.TypeHero)
	hero := game.NewHero(request.UserId, typeHero, request.Name)
	s.saver.Save(ctx, hero)
	childSpan := opentracing.StartSpan("CreateHero -> Save", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	childSpan.LogFields(opentracingLog.Object("Hero", hero))

	result := api.Hero{
		Id:          hero.ID.String(),
		UserId:      hero.UserID,
		TypeHero:    hero.TypeHero.String(),
		Name:        hero.Name,
		Description: hero.Description,
	}

	err := s.kafkaProducer.Send(kafka.Message{
		MessageType: kafka.Create,
		Value:       hero,
	})
	if err != nil {
		s.logger.Info().Msgf("Can not send event to kafka, error: %s", err)
		return nil, err
	}
	s.metrics.IncSuccessCreateHeroCounter()
	return &api.CreateHeroResponse{Hero: &result}, nil
}

func (s *HeroServer) ListHeroes(ctx context.Context, request *api.ListHeroRequest) (*api.ListHeroResponse, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "ListHeroes")
	defer parentSpan.Finish()
	s.logger.Info().Msgf("ListHeroes request: %v", request)

	heroes, err := s.repository.ListHeroes(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	result := make([]*api.Hero, 0, len(heroes))
	for _, hero := range heroes {
		result = append(result, &api.Hero{
			Id:          hero.ID.String(),
			UserId:      hero.UserID,
			TypeHero:    hero.TypeHero.String(),
			Name:        hero.Name,
			Description: hero.Description,
		})
	}
	childSpan := opentracing.StartSpan("ListHeroes -> ListHeroes", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	childSpan.LogFields(opentracingLog.Object("Heroes", result))

	s.metrics.IncSuccessListHeroesCounter()
	return &api.ListHeroResponse{Heroes: result}, nil
}

func (s *HeroServer) DescribeHero(ctx context.Context, request *api.DescribeHeroRequest) (*emptypb.Empty, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "DescribeHero")
	defer parentSpan.Finish()
	s.logger.Info().Msgf("DescribeHero request: %v", request)

	idUuid, err := uuid.Parse(string(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	hero, err := s.repository.DescribeHero(ctx, idUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	childSpan := opentracing.StartSpan("DescribeHero -> DescribeHero", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	childSpan.LogFields(opentracingLog.String("DescribeHero", idUuid.String()))

	err = s.kafkaProducer.Send(kafka.Message{
		MessageType: kafka.Describe,
		Value:       hero,
	})
	if err != nil {
		s.logger.Info().Msgf("Can not send event to kafka, error: %s", err)
		return nil, err
	}
	s.metrics.IncSuccessDescribeHeroCounter()
	return &emptypb.Empty{}, nil
}

func (s *HeroServer) RemoveHero(ctx context.Context, request *api.RemoveHeroRequest) (*emptypb.Empty, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "RemoveHero")
	defer parentSpan.Finish()
	s.logger.Info().Msgf("RemoveHero request: %v", request)

	idUuid, err := uuid.Parse(string(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.repository.RemoveHero(ctx, idUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	childSpan := opentracing.StartSpan("RemoveHero -> RemoveHero", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	childSpan.LogFields(opentracingLog.String("RemoveHero", idUuid.String()))

	s.logger.Info().Msgf("RemoveHero with id=%s removed", string(request.Id))

	err = s.kafkaProducer.Send(kafka.Message{
		MessageType: kafka.Delete,
		Value:       idUuid,
	})
	if err != nil {
		s.logger.Info().Msgf("Can not send event to kafka, error: %s", err)
		return nil, err
	}
	s.metrics.IncSuccessRemoveHeroCounter()
	return &emptypb.Empty{}, nil
}

func (s *HeroServer) UpdateHero(ctx context.Context, request *api.UpdateHeroRequest) (*api.UpdateHeroResponse, error) {
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "UpdateHero")
	defer parentSpan.Finish()
	s.logger.Info().Msgf("UpdateHero request: %v", request)

	idUuid, err := uuid.Parse(string(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	typeHero := game.SearchTypeHeroesEnums(request.TypeHero)

	hero := game.Hero{
		ID:          idUuid,
		UserID:      request.UserId,
		TypeHero:    typeHero,
		Name:        request.Name,
		Description: request.Description,
	}
	err = s.repository.UpdateHero(ctx, hero)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	childSpan := opentracing.StartSpan("UpdateHero -> UpdateHero", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	childSpan.LogFields(opentracingLog.Object("Hero", hero))

	result := api.Hero{
		Id:          hero.ID.String(),
		UserId:      hero.UserID,
		TypeHero:    hero.TypeHero.String(),
		Name:        hero.Name,
		Description: hero.Description,
	}

	err = s.kafkaProducer.Send(kafka.Message{
		MessageType: kafka.Update,
		Value:       hero,
	})
	if err != nil {
		s.logger.Info().Msgf("Can not send event to kafka, error: %s", err)
		return nil, err
	}

	s.metrics.IncSuccessUpdateHeroCounter()
	return &api.UpdateHeroResponse{Hero: &result}, nil
}

func (s *HeroServer) saveHeroesChunks(ctx context.Context, parentSpan opentracing.Span, chunkHeroes []game.Hero) error {
	childSpan := opentracing.StartSpan("MultiCreateHero -> saveHeroesChunks", opentracing.ChildOf(parentSpan.Context()))
	childSpan.LogFields(
		opentracingLog.Int("Heroes count", len(chunkHeroes)),
	)
	defer childSpan.Finish()
	if err := s.repository.AddHeroes(ctx, chunkHeroes); err != nil {
		s.logger.Info().Msgf("Failed add new heroes, error: %s", err)
		childSpan.LogFields(
			opentracingLog.Int("Heroes count failed save", len(chunkHeroes)),
		)
		return err
	}
	return nil
}
