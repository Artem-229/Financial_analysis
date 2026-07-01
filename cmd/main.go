package main

import (
	"context"
	"financial_assistant/internal/app"
)

func main() {
	ctx := context.Background()
	application := app.New(ctx)
	application.Start()
}
