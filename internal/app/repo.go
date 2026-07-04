package app

import (
	"context"
	"financial_assistant/internal/config"
	postgresdb "financial_assistant/internal/infra/postgres"
	"financial_assistant/internal/infra/postgres/repository"
	"fmt"
)

type Repo struct {
	userRepo         *repository.UserRepo
	transactionsRepo *repository.TransactionsRepo
}

func NewRepo(ctx context.Context, creds *config.Credentials) (*Repo, error) {
	pool, err := postgresdb.New(ctx, creds)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	userRepo := repository.NewUserRepo(pool)
	transactionsRepo := repository.NewTransactionsRepo(pool)

	return &Repo{
		userRepo:         userRepo,
		transactionsRepo: transactionsRepo,
	}, nil
}
