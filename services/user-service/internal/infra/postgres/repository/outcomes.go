package repository

import (
	"context"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"

	"github.com/jackc/pgx/v5"
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

	batch := &pgx.Batch{}
	for _, outcome := range outcomes {
		batch.Queue(query, outcome.ID, outcome.UserID, outcome.Analysis)
	}

	results := repo.pool.SendBatch(ctx, batch)
	defer results.Close()

	for range outcomes {
		if _, err := results.Exec(); err != nil {
			return fmt.Errorf("failed to insert outcomes: %w", err)
		}
	}

	return nil
}
