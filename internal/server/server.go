package server

import (
	"context"
	"fmt"
	"github.com/ozonva/ova-game-api/internal/configs"
	"github.com/ozonva/ova-game-api/internal/db"
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/internal/kafka"
	"github.com/ozonva/ova-game-api/internal/logs"
	"github.com/ozonva/ova-game-api/internal/metrics"
	"github.com/ozonva/ova-game-api/internal/repo"
	"github.com/ozonva/ova-game-api/internal/saver"
	"github.com/ozonva/ova-game-api/internal/tracer"
	api "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
)

func Run() error {
	ctx := context.Background()

	logs.InitLogger()
	defer logs.FileLogger.Close()

	configs.LoadConfigs()

	log.Info().Msg("starting server...")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.AppConfig.GrpcPort))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	log.Info().Msg("starting database...")
	pool, err := db.Connect(ctx, configs.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer pool.Close()

	metricClient := metrics.NewApiMetrics("metrics_api", "gRPC_server")

	metrics.RunServer()

	closer, err := tracer.InitTracer(configs.AppConfig.Name, configs.MetricsConfig)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Tracer close error")
		}
	}(closer)

	kafkaClient, err := kafka.NewKafkaProducer(configs.KafkaConfig)
	if err != nil {
		log.Fatal().Msgf("failed connect to kafka, error: %s", err)
	}
	defer func(kafkaClient kafka.KafkaProducer) {
		err := kafkaClient.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Kafka close error")
		}
	}(kafkaClient)

	log.Info().Msgf("start serve port %s", configs.AppConfig.GrpcPort)
	serverGrpc := grpc.NewServer()
	repository := repo.NewHeroRepo(pool)
	fsh := flusher.NewFlusher(configs.AppConfig.SaverChunkSize, repository)
	svr := saver.NewSaver(ctx, configs.AppConfig.SaverChunkSize, fsh, time.Duration(configs.AppConfig.SaverChunkTime)*time.Second)
	api.RegisterHeroApiServer(serverGrpc, NewHeroApiServer(&log.Logger, repository, svr, metricClient, kafkaClient))
	if err := serverGrpc.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}
