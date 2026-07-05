package usecases

import (
	"context"
	"encoding/json"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

type OutcomesRepo interface {
	CreateOutcomes(ctx context.Context, outcomes []entities.Outcome) error
}

type Consumer interface {
	FetchBatch(ctx context.Context, batchSize int, maxWait time.Duration) ([]kafka.Message, error)
	CommitMessages(ctx context.Context, messages ...kafka.Message) error
}

type OutcomesUsecases struct {
	KafkaReader Consumer
	Repo        OutcomesRepo
	BatchSize   int
	MaxWait     time.Duration
}

type OutcomesUsecasesDeps struct {
	KafkaReader Consumer
	Repo        OutcomesRepo
	BatchSize   int
	MaxWait     time.Duration
}

func NewOutcomesUsecases(deps OutcomesUsecasesDeps) *OutcomesUsecases {
	return &OutcomesUsecases{
		KafkaReader: deps.KafkaReader,
		Repo:        deps.Repo,
		BatchSize:   deps.BatchSize,
		MaxWait:     deps.MaxWait,
	}
}

// Run consumes batches of outcome events produced by the analyzer-service
// until ctx is cancelled.
func (t *OutcomesUsecases) Run(ctx context.Context) {
	for {
		if err := t.consumeBatch(ctx); err != nil {
			if ctx.Err() != nil {
				return
			}
			slog.Error("consume outcomes batch", "error", err)
		}
	}
}

func (t *OutcomesUsecases) consumeBatch(ctx context.Context) error {
	messages, err := t.KafkaReader.FetchBatch(ctx, t.BatchSize, t.MaxWait)
	if err != nil {
		return fmt.Errorf("could not fetch batch: %w", err)
	}

	// nothing arrived within MaxWait - nothing to do, loop again.
	if len(messages) == 0 {
		return nil
	}

	outcomes := make([]entities.Outcome, 0, len(messages))
	for _, message := range messages {
		var outcome entities.Outcome
		if err := json.Unmarshal(message.Value, &outcome); err != nil {
			slog.Error("skipping malformed outcome message", "error", err)
			continue
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) > 0 {
		if err := t.Repo.CreateOutcomes(ctx, outcomes); err != nil {
			return fmt.Errorf("could not create outcomes: %w", err)
		}
	}

	// Commit only after the whole batch has been processed successfully, so a
	// crash mid-processing leaves the batch uncommitted and it gets
	// re-delivered instead of silently lost.
	if err := t.KafkaReader.CommitMessages(ctx, messages...); err != nil {
		return fmt.Errorf("could not commit messages: %w", err)
	}

	return nil
}
