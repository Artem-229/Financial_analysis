package app

import (
	"financial_assistant/services/user-service/internal/config"
	"financial_assistant/services/user-service/internal/infra/kafka/consumer"
	"financial_assistant/services/user-service/internal/infra/kafka/producer"
)

type Kafka struct {
	Producer *producer.Producer
	Consumer *consumer.Consumer
}

func NewKafka(conf *config.Config) *Kafka {
	return &Kafka{
		Producer: producer.NewProducer(conf.Kafka.Brokers, conf.Kafka.UpstreamTopic),
		Consumer: consumer.NewConsumer(conf.Kafka.Brokers, conf.Kafka.HandleTopic),
	}
}
