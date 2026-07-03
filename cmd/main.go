package main

import (
	"context"
	"financial_assistant/internal/app"
	"financial_assistant/internal/config"
	"fmt"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %w", err))
	}

	creds, err := config.ReadCredentials()
	if err != nil {
		panic(fmt.Errorf("error reading credentials: %w", err))
	}

	ctx := context.Background()
	application := app.New(ctx, conf, creds)
	if err := application.Start(); err != nil {
		panic(fmt.Errorf("error starting application: %w", err))
	}
}
