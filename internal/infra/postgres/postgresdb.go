package postgresdb

import (
	"context"
	"financial_assistant/internal/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func New(ctx context.Context, creds *config.Credentials) (*pgxpool.Pool, error) {
	pgx, err := pgxpool.New(ctx, creds.Postgres.Connstr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	db := stdlib.OpenDBFromPool(pgx)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("error setting postgres dialect: %w", err)
	}
	if err := goose.Up(db, "/migrations"); err != nil {
		return nil, fmt.Errorf("error upgrading migrations: %w", err)
	}

	return pgx, nil
}
