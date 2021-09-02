package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozonva/ova-game-api/internal/configs"
	"github.com/rs/zerolog/log"
)

// KafkaProducer - interface for work with Kafka
type KafkaProducer interface {
	Send(message Message) error
	Close() error
}

type kafkaProducer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

// NewKafkaProducer - creates new Producer for work with Kafka
func NewKafkaProducer(configuration *configs.Kafka) (KafkaProducer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true

	brokers := make([]string, 0, len(configuration.Brokers))
	for _, broker := range configuration.Brokers {
		brokers = append(brokers, fmt.Sprintf("%s:%s", broker.Host, broker.Port))
	}

	syncProducer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Kafka producer: failed to create")
		return nil, err
	}

	return &kafkaProducer{
		topic:        configuration.Topic,
		syncProducer: syncProducer,
	}, nil
}

// Send - Send new message to Kafka
func (p *kafkaProducer) Send(message Message) error {
	jsonMes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = p.syncProducer.SendMessage(
		&sarama.ProducerMessage{
			Topic:     p.topic,
			Partition: -1,
			Key:       sarama.StringEncoder(p.topic),
			Value:     sarama.StringEncoder(jsonMes),
		})
	return err
}

func (p *kafkaProducer) Close() error {
	return p.syncProducer.Close()
}
