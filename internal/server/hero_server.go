package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/ozonva/ova-game-api/internal/repo"
	"github.com/ozonva/ova-game-api/internal/saver"
	"github.com/ozonva/ova-game-api/pkg/game"
	api "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HeroServer struct {
	api.UnimplementedHeroApiServer
	logger     *zerolog.Logger
	repository repo.HeroRepo
	saver      saver.Saver
}

func NewHeroApiServer(logger *zerolog.Logger, repo repo.HeroRepo, saver saver.Saver) api.HeroApiServer {
	return &HeroServer{
		UnimplementedHeroApiServer: api.UnimplementedHeroApiServer{},
		logger:                     logger,
		repository:                 repo,
		saver:                      saver,
	}
}

func (s *HeroServer) CreateHero(context context.Context, request *api.CreateHeroRequest) (*api.CreateHeroResponse, error) {
	s.logger.Info().Msgf("CreateHero request: %v", request)

	typeHero := game.SearchTypeHeroesEnums(request.TypeHero)
	hero := game.NewHero(request.UserId, typeHero, request.Name)
	s.saver.Save(hero)

	result := api.Hero{
		Id:       hero.ID.String(),
		UserId:   hero.UserID,
		TypeHero: hero.TypeHero.String(),
		Name:     hero.Name,
	}

	return &api.CreateHeroResponse{Hero: &result}, nil
}

func (s *HeroServer) ListHeroes(context context.Context, request *api.ListHeroRequest) (*api.ListHeroResponse, error) {
	s.logger.Info().Msgf("ListHeroes request: %v", request)

	heroes, err := s.repository.ListHeroes(request.Limit, request.Offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	result := make([]*api.Hero, 0, len(heroes))
	for _, hero := range heroes {
		result = append(result, &api.Hero{
			Id:       hero.ID.String(),
			UserId:   hero.UserID,
			TypeHero: hero.TypeHero.String(),
			Name:     hero.Name,
		})
	}

	return &api.ListHeroResponse{Heroes: result}, nil
}

func (s *HeroServer) DescribeHero(context context.Context, request *api.DescribeHeroRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("DescribeHero request: %v", request)

	idUuid, err := uuid.Parse(string(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	_, err = s.repository.DescribeHero(idUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *HeroServer) RemoveHero(context context.Context, request *api.RemoveHeroRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveHero request: %v", request)

	idUuid, err := uuid.Parse(string(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.repository.RemoveHero(idUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.logger.Info().Msgf("RemoveHero with id=%s removed", string(request.Id))

	return &emptypb.Empty{}, nil
}
