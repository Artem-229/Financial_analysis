package app

import (
	"financial_assistant/services/analyzer-service/internal/config"
	"financial_assistant/services/analyzer-service/internal/infra/kafka/consumer"
	"financial_assistant/services/analyzer-service/internal/infra/kafka/producer"
)

type Kafka struct {
	Producer *producer.Producer
	Consumer *consumer.Consumer
}

func NewKafka(conf *config.Config) *Kafka {
	return &Kafka{
		Consumer: consumer.NewConsumer(conf.Kafka.Brokers, conf.Kafka.UpstreamTopic, conf.Kafka.ConsumerGroupID),
		Producer: producer.NewProducer(conf.Kafka.Brokers, conf.Kafka.HandleTopic),
	}
}

func (k *Kafka) Close() error {
	if err := k.Producer.Close(); err != nil {
		return err
	}
	return k.Consumer.Close()
}
