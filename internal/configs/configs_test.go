package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigsLoad(t *testing.T) {
	appConfig := App{}
	databaseConfig := Database{}
	kafkaConfig := Kafka{
		Brokers: []*BrokerKafka{{}},
	}
	metricsConfig := Metrics{
		Prometheus: Prometheus{},
		Jaeger:     Jaeger{},
	}

	LoadConfigs()
	assert.Equal(t, appConfig, *AppConfig)
	assert.Equal(t, databaseConfig, *DatabaseConfig)
	assert.Equal(t, kafkaConfig, *KafkaConfig)
	assert.Equal(t, metricsConfig, *MetricsConfig)
}
