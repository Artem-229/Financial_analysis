package repository

import (
	"context"
	"financial_assistant/internal/entities"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		pool: pool,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO users (id, name, login, password_hash) VALUES ($1, $2, $3, $4)`
	if _, err := u.pool.Exec(ctx, query, user.Id, user.Username, user.Login, user.Password); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserRepo) GetUserByLoginAndPassword(ctx context.Context, login string) (*entities.User, error) {
	query := `SELECT * FROM users WHERE login = $1`

	res := u.pool.QueryRow(ctx, query, login)

	var user entities.User
	if err := res.Scan(
		&user.Id,
		&user.Username,
		&user.Login,
		&user.Password,
	); err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &user, nil
}
