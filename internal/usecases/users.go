package usecases

import (
	"context"
	"financial_assistant/internal/entities"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entities.User) error
}
type Usecases struct {
	UserRepo UserRepo
}

type UsecasesDeps struct {
	UserRepo UserRepo
}

func NewUsecases(deps UsecasesDeps) *Usecases {
	return &Usecases{
		deps.UserRepo,
	}
}

func (u Usecases) CreateUser(ctx context.Context, user *entities.User) error {
	return u.UserRepo.CreateUser(ctx, user)
}
