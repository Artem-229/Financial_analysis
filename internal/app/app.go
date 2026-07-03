package app

import (
	"context"
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/rest"
	"fmt"
	"log"
	"log/slog"
)

type App struct {
	rest     *rest.Server
	repo     *Repo
	usecases *Usecases
	logger   slog.Logger
}

func New(ctx context.Context, conf *config.Config, creds *config.Credentials) *App {
	repo, err := NewRepo(ctx, creds)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	usecases, err := NewUsecases(ctx, repo)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	server := rest.NewServer(&rest.ServerDeps{
		Config:   conf,
		Usecases: usecases.User,
	})

	return &App{
		rest:     server,
		repo:     repo,
		usecases: usecases,
		logger:   slog.Logger{},
	}
}

func (app *App) Start() error {
	if err := app.rest.Start(); err != nil {
		return fmt.Errorf("start application: %w", err)
	}
	return nil
}
