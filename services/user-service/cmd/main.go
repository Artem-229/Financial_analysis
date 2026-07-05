package main

import (
	"context"
	"financial_assistant/services/user-service/internal/app"
	config2 "financial_assistant/services/user-service/internal/config"
	"fmt"
)

func main() {
	conf, err := config2.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %w", err))
	}

	creds, err := config2.ReadCredentials()
	if err != nil {
		panic(fmt.Errorf("error reading credentials: %w", err))
	}

	ctx := context.Background()
	application := app.New(ctx, conf, creds)
	if err := application.Start(); err != nil {
		panic(fmt.Errorf("error starting application: %w", err))
	}
}
