package app

import (
	"context"
	"financial_assistant/services/analyzer-service/internal/config"
	"log"
	"log/slog"
)

type App struct {
	ctx      context.Context
	usecases *Usecases
	kafka    *Kafka
}

func New(ctx context.Context, conf *config.Config) *App {
	kafka := NewKafka(conf)

	usecases, err := NewUsecases(ctx, kafka, conf)
	if err != nil {
		log.Fatalf("error initializing usecases: %v", err)
	}

	go usecases.Analyzer.Run(ctx)

	return &App{
		ctx:      ctx,
		usecases: usecases,
		kafka:    kafka,
	}
}

func (app *App) Start() error {
	<-app.ctx.Done()
	slog.Info("shutting down analyzer-service")
	return app.kafka.Close()
}
