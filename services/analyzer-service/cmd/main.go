package main

import (
	"context"
	"financial_assistant/services/analyzer-service/internal/app"
	"financial_assistant/services/analyzer-service/internal/config"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	conf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	application := app.New(ctx, conf)
	if err := application.Start(); err != nil {
		panic(fmt.Errorf("error starting application: %w", err))
	}
}
