package app

import (
	"context"
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/rest"
	"fmt"
	"log/slog"
)

type App struct {
	rest   rest.Server
	logger slog.Logger
}

func New(ctx context.Context, conf *config.Config) *App {
	server := rest.NewServer(&rest.ServerDeps{
		Config: conf,
	})

	return &App{
		rest: *server,
	}
}

func (app *App) Start() error {
	if err := app.rest.Start(); err != nil {
		return fmt.Errorf("start application: %w", err)
	}
	return nil
}
