package metrics

import (
	"fmt"
	"net/http"

	"github.com/ozonva/ova-game-api/internal/configs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s", configs.MetricsConfig.Prometheus.Path), promhttp.Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.MetricsConfig.Prometheus.Port),
		Handler: mux,
	}

	log.Info().Msg("starting server metrics...")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()
}
