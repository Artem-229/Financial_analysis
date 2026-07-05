package app

import (
	"context"
	"financial_assistant/services/analyzer-service/internal/config"
	"financial_assistant/services/analyzer-service/internal/usecases"
)

type Usecases struct {
	Analyzer *usecases.AnalyzerUsecases
}

func NewUsecases(ctx context.Context, kafka *Kafka, conf *config.Config) (*Usecases, error) {
	analyzerUsecases := usecases.NewAnalyzerUsecases(usecases.AnalyzerUsecasesDeps{
		Consumer:  kafka.Consumer,
		Producer:  kafka.Producer,
		LLM:       usecases.NewMockLLM(),
		BatchSize: conf.Kafka.BatchSize,
		MaxWait:   conf.Kafka.MaxWait,
	})

	return &Usecases{
		Analyzer: analyzerUsecases,
	}, nil
}
