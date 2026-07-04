package app

import (
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/kafka/producer"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Producer *producer.Producer
	consumer *kafka.Reader
}

func NewKafka(conf *config.Config) *Kafka {
	return &Kafka{
		Producer: producer.NewProducer(conf.Kafka.Brokers, conf.Kafka.UpstreamTopic),
		consumer: kafka.NewReader(kafka.ReaderConfig{
			Brokers: conf.Kafka.Brokers,
			Topic:   conf.Kafka.HandleTopic,
		}),
	}
}
