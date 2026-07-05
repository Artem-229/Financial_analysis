package app

import (
	"context"
	"financial_assistant/services/user-service/internal/config"
	"financial_assistant/services/user-service/internal/usecases"
)

type Usecases struct {
	User         *usecases.UserUsecases
	Transactions *usecases.TransactionsUsecases
	Outcomes     *usecases.OutcomesUsecases
}

func NewUsecases(ctx context.Context, repo *Repo, kafka *Kafka, conf *config.Config, creds *config.Credentials) (*Usecases, error) {
	userUsecase := usecases.NewUserUsecases(usecases.UserDeps{
		UserRepo: repo.userRepo,
		Creds:    creds,
	})

	transcationsUsecases := usecases.NewTranscationsUsecases(usecases.TransactionsDeps{
		Repo:        repo.transactionsRepo,
		KafkaWriter: kafka.Producer,
	})

	outcomesUsecases := usecases.NewOutcomesUsecases(usecases.OutcomesUsecasesDeps{
		Repo:        repo.outcomesRepo,
		KafkaReader: kafka.Consumer,
		BatchSize:   conf.Kafka.BatchSize,
		MaxWait:     conf.Kafka.MaxWait,
	})

	return &Usecases{
		User:         userUsecase,
		Transactions: transcationsUsecases,
		Outcomes:     outcomesUsecases,
	}, nil
}
