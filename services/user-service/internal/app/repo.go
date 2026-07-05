package app

import (
	"context"
	"financial_assistant/services/user-service/internal/config"
	"financial_assistant/services/user-service/internal/infra/postgres"
	"financial_assistant/services/user-service/internal/infra/postgres/repository"
	"fmt"
)

type Repo struct {
	userRepo         *repository.UserRepo
	transactionsRepo *repository.TransactionsRepo
	outcomesRepo     *repository.OutcomesRepo
}

func NewRepo(ctx context.Context, creds *config.Credentials) (*Repo, error) {
	pool, err := postgresdb.New(ctx, creds)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	userRepo := repository.NewUserRepo(pool)
	transactionsRepo := repository.NewTransactionsRepo(pool)
	outcomesRepo := repository.NewOutcomesRepo(pool)

	return &Repo{
		userRepo:         userRepo,
		transactionsRepo: transactionsRepo,
		outcomesRepo:     outcomesRepo,
	}, nil
}
