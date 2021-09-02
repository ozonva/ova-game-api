package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigsLoad(t *testing.T) {
	appConfig := App{
		Name:           "",
		Environment:    "",
		Debug:          false,
		GrpcPort:       "",
		SaverChunkSize: 0,
		SaverChunkTime: 0,
	}
	databaseConfig := Database{
		DbName:         "",
		Host:           "",
		Port:           "",
		Username:       "",
		Password:       "",
		PoolMaxConnect: 0,
	}
	kafkaConfig := Kafka{
		Topic:   "",
		Brokers: []*BrokerKafka{{Host: "", Port: ""}},
	}
	metricsConfig := Metrics{
		Prometheus: Prometheus{Host: "", Port: "", Path: ""},
		Jaeger:     Jaeger{Host: "", Port: ""},
	}

	LoadConfigs()
	assert.Equal(t, appConfig, *AppConfig)
	assert.Equal(t, databaseConfig, *DatabaseConfig)
	assert.Equal(t, kafkaConfig, *KafkaConfig)
	assert.Equal(t, metricsConfig, *MetricsConfig)
}
