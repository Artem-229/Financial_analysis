package app

import (
	"context"
	"financial_assistant/internal/config"
	"financial_assistant/internal/usecases"
)

type Usecases struct {
	User *usecases.Usecases
}

func NewUsecases(ctx context.Context, repo *Repo, creds *config.Credentials) (*Usecases, error) {
	userUsecase := usecases.NewUsecases(usecases.Deps{
		UserRepo: repo.userRepo,
		Creds:    creds,
	})
	return &Usecases{
		User: userUsecase,
	}, nil
}
