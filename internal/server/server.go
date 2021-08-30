package server

import (
	api "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

const (
	grpcPort = ":82"
)

func Run() error {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	log := zerolog.New(output).With().Timestamp().Logger()
	log.Info().Msgf("start serve port %s", grpcPort)

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterHeroApiServer(s, NewHeroApiServer(&log))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}
