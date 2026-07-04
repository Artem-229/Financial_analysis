package app

import (
	"context"
	"financial_assistant/internal/config"
	"financial_assistant/internal/usecases"
)

type Usecases struct {
	User         *usecases.UserUsecases
	Transactions *usecases.TransactionsUsecases
}

func NewUsecases(ctx context.Context, repo *Repo, creds *config.Credentials) (*Usecases, error) {
	userUsecase := usecases.NewUserUsecases(usecases.UserDeps{
		UserRepo: repo.userRepo,
		Creds:    creds,
	})
	transcationsUsecases := usecases.NewTranscationsUsecases(usecases.TransactionsDeps{
		Repo: repo.transactionsRepo,
	})

	return &Usecases{
		User:         userUsecase,
		Transactions: transcationsUsecases,
	}, nil
}
