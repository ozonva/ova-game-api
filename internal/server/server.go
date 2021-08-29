package server

import (
	"context"
	"github.com/ozonva/ova-game-api/internal/configs"
	"github.com/ozonva/ova-game-api/internal/db"
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/internal/repo"
	"github.com/ozonva/ova-game-api/internal/saver"
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
	ctx := context.Background()

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	configs.LoadConfigs()

	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Msg("starting server...")
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	log.Info().Msg("starting database...")
	pool, err := db.Connect(ctx, configs.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer pool.Close()

	log.Info().Msgf("start serve port %s", grpcPort)
	s := grpc.NewServer()
	repository := repo.NewHeroRepo(ctx, pool)
	fsh := flusher.NewFlusher(5, repository)
	svr := saver.NewSaver(5, fsh, 2)
	api.RegisterHeroApiServer(s, NewHeroApiServer(&log, repository, svr))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}
