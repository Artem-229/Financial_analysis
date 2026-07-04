package usecases

import (
	"context"
	"financial_assistant/internal/entities"
)

type Repo interface {
	Upsert(ctx context.Context, transaction *entities.Transaction) error
	GetByID(ctx context.Context, id string, userID string) (*entities.Transaction, error)
	List(ctx context.Context, userID string) ([]entities.Transaction, error)
	Delete(ctx context.Context, id string, userID string) error
}

type TransactionsUsecases struct {
	Repo Repo
}

type TransactionsDeps struct {
	Repo Repo
}

func NewTranscationsUsecases(deps TransactionsDeps) *TransactionsUsecases {
	return &TransactionsUsecases{
		Repo: deps.Repo,
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
