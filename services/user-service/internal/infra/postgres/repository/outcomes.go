package repository

import (
	"context"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OutcomesRepo struct {
	pool *pgxpool.Pool
}

func NewOutcomesRepo(pool *pgxpool.Pool) *OutcomesRepo {
	return &OutcomesRepo{
		pool: pool,
	}
}

func (repo *OutcomesRepo) CreateOutcomes(ctx context.Context, outcomes []entities.Outcome) error {
	query := `INSERT INTO outcomes (id, user_id, analysis) VALUES ($1, $2, $3)`

	_, err := repo.pool.Exec(ctx, query, outcomes)
	if err != nil {
		return fmt.Errorf("failed to insert outcomes: %w", err)
	}

	return nil
}
