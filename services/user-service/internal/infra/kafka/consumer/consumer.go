package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	return &Consumer{kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
	})}
}

func (consumer *Consumer) Consume(ctx context.Context) ([]*kafka.Message, error) {
	messages := make([]*kafka.Message, 0)

	for len(messages) < 10 {
		message, err := consumer.reader.FetchMessage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to read message: %w", err)
		}

		messages = append(messages, &message)
	}
	if err := consumer.reader.CommitMessages(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit message: %w", err)
	}
	return messages, nil
}
