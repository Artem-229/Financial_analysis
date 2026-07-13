package usecases

import (
	"context"
	"encoding/json"
	"financial_assistant/services/analyzer-service/internal/entities"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	FetchBatch(ctx context.Context, batchSize int, maxWait time.Duration) ([]kafka.Message, error)
	CommitMessages(ctx context.Context, messages ...kafka.Message) error
}

type Producer interface {
	Produce(ctx context.Context, events [][]byte) error
}

type LLM interface {
	Analyze(ctx context.Context, transactions map[uuid.UUID]map[string]entities.SortedTransactions) ([]entities.Outcome, error)
}

type AnalyzerUsecases struct {
	Consumer  Consumer
	Producer  Producer
	LLM       LLM
	BatchSize int
	MaxWait   time.Duration
}

type AnalyzerUsecasesDeps struct {
	Consumer  Consumer
	Producer  Producer
	LLM       LLM
	BatchSize int
	MaxWait   time.Duration
}

func NewAnalyzerUsecases(deps AnalyzerUsecasesDeps) *AnalyzerUsecases {
	return &AnalyzerUsecases{
		Consumer:  deps.Consumer,
		Producer:  deps.Producer,
		LLM:       deps.LLM,
		BatchSize: deps.BatchSize,
		MaxWait:   deps.MaxWait,
	}
}

func (a *AnalyzerUsecases) Run(ctx context.Context) {
	for {
		if err := a.processBatch(ctx); err != nil {
			if ctx.Err() != nil {
				return
			}
			slog.Error("process transactions batch", "error", err)
		}
	}
}

func (a *AnalyzerUsecases) processBatch(ctx context.Context) error {
	messages, err := a.Consumer.FetchBatch(ctx, a.BatchSize, a.MaxWait)
	if err != nil {
		return fmt.Errorf("could not fetch batch: %w", err)
	}

	if len(messages) == 0 {
		return nil
	}

	transactions := make([]entities.Transaction, 0, len(messages))
	for _, message := range messages {
		var transaction entities.Transaction
		if err := json.Unmarshal(message.Value, &transaction); err != nil {
			slog.Error("skipping malformed transaction message", "error", err)
			continue
		}
		transactions = append(transactions, transaction)
	}

	if len(transactions) > 0 {
		outcomes, err := a.LLM.Analyze(ctx, sort(transactions))
		if err != nil {
			return fmt.Errorf("could not analyze batch: %w", err)
		}

		events := make([][]byte, len(outcomes))
		for i, outcome := range outcomes {
			data, err := json.Marshal(outcome)
			if err != nil {
				return fmt.Errorf("could not marshal outcome %s: %w", outcome.ID, err)
			}
			events[i] = data
		}

		if err := a.Producer.Produce(ctx, events); err != nil {
			return fmt.Errorf("could not produce outcomes: %w", err)
		}
	}

	if err := a.Consumer.CommitMessages(ctx, messages...); err != nil {
		return fmt.Errorf("could not commit messages: %w", err)
	}

	return nil
}

func sort(transactions []entities.Transaction) map[uuid.UUID]map[string]entities.SortedTransactions {
	result := make(map[uuid.UUID]map[string]entities.SortedTransactions)
	for _, transaction := range transactions {
		byClass, ok := result[transaction.UserID]
		if !ok {
			byClass = make(map[string]entities.SortedTransactions)
			result[transaction.UserID] = byClass
		}

		temp := byClass[transaction.Class]
		temp.Price += transaction.Price
		temp.Items = append(temp.Items, transaction.Item)
		temp.Notes = append(temp.Notes, transaction.Note)
		byClass[transaction.Class] = temp
	}

	return result
}
