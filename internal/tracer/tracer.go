package tracer

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-game-api/internal/configs"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func InitTracer(serviceName string, metricsConfig *configs.Metrics) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%s", metricsConfig.Jaeger.Host, metricsConfig.Jaeger.Port),
		},
	}

	jLogger := jaegerlog.StdLogger

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
	)
	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}
