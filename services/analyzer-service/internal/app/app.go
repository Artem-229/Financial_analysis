package app

import (
	"context"
	"financial_assistant/services/analyzer-service/internal/config"
	"log"
	"log/slog"
)

type App struct {
	usecases *Usecases
	kafka    *Kafka
	logger   slog.Logger
}

func New(ctx context.Context, conf *config.Config) *App {
	kafka := NewKafka(conf)

	usecases, err := NewUsecases(ctx, kafka)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	return &App{
		usecases: usecases,
		kafka:    kafka,
		logger:   slog.Logger{},
	}
}

func (app *App) Start() error {
	return nil
}
