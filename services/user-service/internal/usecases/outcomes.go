package usecases

import (
	"context"
	"encoding/json"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type OutcomesRepo interface {
	CreateOutcomes(ctx context.Context, outcomes []entities.Outcome) error
}

type Consumer interface {
	Consume(ctx context.Context) ([]*kafka.Message, error)
}

type OutcomesUsecases struct {
	KafkaReader Consumer
	Repo        OutcomesRepo
}

type OutcomesUsecasesDeps struct {
	KafkaReader Consumer
	Repo        OutcomesRepo
}

func NewOutcomesUsecases(deps OutcomesUsecasesDeps) *OutcomesUsecases {
	return &OutcomesUsecases{
		KafkaReader: deps.KafkaReader,
		Repo:        deps.Repo,
	}
}

func (t *OutcomesUsecases) Consume(ctx context.Context) error {
	messages, err := t.KafkaReader.Consume(ctx)
	if err != nil {
		return fmt.Errorf("could not consume messages: %w", err)
	}

	if len(messages) == 0 {
		return nil
	}

	outcomes := make([]entities.Outcome, 0)

	for _, message := range messages {
		var outcome entities.Outcome
		value := message.Value
		if err := json.Unmarshal(value, &outcome); err != nil {
			return fmt.Errorf("could not unmarshal value: %w", err)
		}
		outcomes = append(outcomes, outcome)
	}

	if err := t.Repo.CreateOutcomes(ctx, outcomes); err != nil {
		return fmt.Errorf("could not create outcomes: %w", err)
	}

	return nil
}
