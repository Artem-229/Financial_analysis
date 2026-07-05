package app

import (
	"context"
	"financial_assistant/services/user-service/internal/config"
	"financial_assistant/services/user-service/internal/infra/rest"
	"fmt"
	"log"
	"log/slog"

	"github.com/robfig/cron/v3"
)

type App struct {
	rest     *rest.Server
	repo     *Repo
	usecases *Usecases
	kafka    *Kafka
	logger   slog.Logger
}

func New(ctx context.Context, conf *config.Config, creds *config.Credentials) *App {
	kafka := NewKafka(conf)

	repo, err := NewRepo(ctx, creds)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	usecases, err := NewUsecases(ctx, repo, kafka, conf, creds)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	scheduler := cron.New(cron.WithSeconds())
	if _, err := scheduler.AddFunc(conf.Cron.Processor.Spec, func() {
		if err := usecases.Transactions.Process(ctx); err != nil {
			log.Printf("process pending transactions: %v", err)
		}
	}); err != nil {
		log.Fatalf("error scheduling transaction processor: %v", err)
	}
	scheduler.Start()

	go usecases.Outcomes.Run(ctx)

	server := rest.NewServer(&rest.ServerDeps{
		Config:       conf,
		Creds:        creds,
		Usecases:     usecases.User,
		Transactions: usecases.Transactions,
	})

	return &App{
		rest:     server,
		repo:     repo,
		usecases: usecases,
		kafka:    kafka,
		logger:   slog.Logger{},
	}
}

func (app *App) Start() error {
	if err := app.rest.Start(); err != nil {
		return fmt.Errorf("start application: %w", err)
	}
	return nil
}
