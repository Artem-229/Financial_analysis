package consumer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

// FetchBatch collects up to batchSize messages, waiting at most maxWait for
// the batch to fill up. It returns as soon as batchSize messages are collected,
// or once maxWait elapses since the call started, whichever happens first -
// so callers never block indefinitely on low traffic. The returned batch can
// be smaller than batchSize (even empty) when maxWait runs out first.
func (c *Consumer) FetchBatch(ctx context.Context, batchSize int, maxWait time.Duration) ([]kafka.Message, error) {
	batchCtx, cancel := context.WithTimeout(ctx, maxWait)
	defer cancel()

	messages := make([]kafka.Message, 0, batchSize)
	for len(messages) < batchSize {
		message, err := c.reader.FetchMessage(batchCtx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}
			return messages, fmt.Errorf("failed to fetch message: %w", err)
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (c *Consumer) CommitMessages(ctx context.Context, messages ...kafka.Message) error {
	if err := c.reader.CommitMessages(ctx, messages...); err != nil {
		return fmt.Errorf("failed to commit messages: %w", err)
	}

	return nil
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
