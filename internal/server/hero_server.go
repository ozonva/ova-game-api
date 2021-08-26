package server

import (
	"context"
	api "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HeroServer struct {
	api.UnimplementedHeroApiServer
	logger *zerolog.Logger
}

func NewHeroApiServer(logger *zerolog.Logger) api.HeroApiServer {
	return &HeroServer{
		UnimplementedHeroApiServer: api.UnimplementedHeroApiServer{},
		logger:                     logger,
	}
}

func (s *HeroServer) CreateHero(context context.Context, request *api.CreateHeroRequest) (*api.CreateGameResponse, error) {
	s.logger.Info().Msgf("CreateHero request: %v", request)
	return s.UnimplementedHeroApiServer.CreateHero(context, request)
}

func (s *HeroServer) ListHeroes(context context.Context, request *api.ListHeroRequest) (*api.ListHeroResponse, error) {
	s.logger.Info().Msgf("ListHeroes request: %v", request)
	return s.UnimplementedHeroApiServer.ListHeroes(context, request)
}

func (s *HeroServer) DescribeHero(context context.Context, request *api.DescribeHeroRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("DescribeHero request: %v", request)
	return s.UnimplementedHeroApiServer.DescribeHero(context, request)
}

func (s *HeroServer) RemoveHero(context context.Context, request *api.RemoveHeroRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveHero request: %v", request)
	return s.UnimplementedHeroApiServer.RemoveHero(context, request)
}
