package usecases

import (
	"context"
	"financial_assistant/services/user-service/internal/config"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByLoginAndPassword(ctx context.Context, login string) (*entities.User, error)
}
type UserUsecases struct {
	UserRepo UserRepo
	Creds    *config.Credentials
}

type UserDeps struct {
	UserRepo UserRepo
	Creds    *config.Credentials
}

func NewUserUsecases(deps UserDeps) *UserUsecases {
	return &UserUsecases{
		UserRepo: deps.UserRepo,
		Creds:    deps.Creds,
	}
}

func (u UserUsecases) CreateUser(ctx context.Context, user *entities.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	user.Password = hash

	return u.UserRepo.CreateUser(ctx, user)
}

func (u UserUsecases) LoginUser(ctx context.Context, login string, password string) (string, error) {
	rawUser, err := u.UserRepo.GetUserByLoginAndPassword(ctx, login)
	if err != nil {
		return "", fmt.Errorf("could not get user: %w", err)
	}

	if !CheckPasswordHash(password, rawUser.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := CreateJWTToken(rawUser.Id.String(), u.Creds.JWT.Secret)
	if err != nil {
		return "", fmt.Errorf("could not create token: %w", err)
	}

	return token, nil
}
