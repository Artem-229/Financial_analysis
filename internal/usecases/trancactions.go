package usecases

import (
	"context"
	"encoding/json"
	"financial_assistant/internal/entities"
	"fmt"

	"github.com/google/uuid"
)

type Repo interface {
	Upsert(ctx context.Context, transaction *entities.Transaction) error
	GetByID(ctx context.Context, id string, userID string) (*entities.Transaction, error)
	List(ctx context.Context, userID string) ([]entities.Transaction, error)
	Delete(ctx context.Context, id string, userID string) error
	GetPendingTransactions(ctx context.Context) ([]entities.Transaction, error)
	MarkTransactionsProcessed(ctx context.Context, ids []uuid.UUID, success bool) error
}

type Producer interface {
	Produce(ctx context.Context, events [][]byte) error
}

type TransactionsUsecases struct {
	Repo        Repo
	KafkaWriter Producer
}

type TransactionsDeps struct {
	Repo        Repo
	KafkaWriter Producer
}

func NewTranscationsUsecases(deps TransactionsDeps) *TransactionsUsecases {
	return &TransactionsUsecases{
		Repo:        deps.Repo,
		KafkaWriter: deps.KafkaWriter,
	}
}

func (t *TransactionsUsecases) UpsertTransaction(ctx context.Context, transaction *entities.Transaction) error {
	return t.Repo.Upsert(ctx, transaction)
}

func (t *TransactionsUsecases) GetTransaction(ctx context.Context, id string, userID string) (*entities.Transaction, error) {
	return t.Repo.GetByID(ctx, id, userID)
}

func (t *TransactionsUsecases) ListTransactions(ctx context.Context, userID string) ([]entities.Transaction, error) {
	return t.Repo.List(ctx, userID)
}

func (t *TransactionsUsecases) DeleteTransaction(ctx context.Context, id string, userID string) error {
	return t.Repo.Delete(ctx, id, userID)
}

func (t *TransactionsUsecases) Process(ctx context.Context) error {
	transactions, err := t.Repo.GetPendingTransactions(ctx)
	if err != nil {
		return fmt.Errorf("could not get pending transactions: %w", err)
	}

	if len(transactions) == 0 {
		return nil
	}

	ids := make([]uuid.UUID, len(transactions))
	events := make([][]byte, len(transactions))

	for i, transaction := range transactions {
		ids[i] = transaction.ID

		data, err := json.Marshal(transaction)
		if err != nil {
			return fmt.Errorf("could not marshal transaction %s: %w", transaction.ID, err)
		}
		events[i] = data
	}

	if err := t.KafkaWriter.Produce(ctx, events); err != nil {
		if markErr := t.Repo.MarkTransactionsProcessed(ctx, ids, false); markErr != nil {
			return fmt.Errorf("could not produce events: %v; could not revert status: %w", err, markErr)
		}
		return fmt.Errorf("could not produce events: %w", err)
	}

	return t.Repo.MarkTransactionsProcessed(ctx, ids, true)
}
