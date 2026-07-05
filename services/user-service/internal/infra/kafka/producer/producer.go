package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Produce(ctx context.Context, events [][]byte) error {
	messages := make([]kafka.Message, len(events))

	for i, e := range events {
		messages[i] = kafka.Message{
			Value: e,
		}
	}

	return p.writer.WriteMessages(ctx, messages...)
}
