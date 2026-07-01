package app

import (
	"context"
	"financial_assistant/internal/infra/rest"
)

type App struct {
	rest rest.Server
}

func New(ctx context.Context) *App {
	server := rest.NewServer()
	return &App{
		rest: *server,
	}
}

func (app *App) Start() {
	app.rest.Engine.Run("localhost:8080")
}
