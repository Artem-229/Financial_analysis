package app

import (
	"context"
	"financial_assistant/internal/usecases"
)

type Usecases struct {
	User *usecases.Usecases
}

func NewUsecases(ctx context.Context, repo *Repo) (*Usecases, error) {
	userUsecase := usecases.NewUsecases(usecases.UsecasesDeps{
		UserRepo: repo.userRepo,
	})
	return &Usecases{
		User: userUsecase,
	}, nil
}
