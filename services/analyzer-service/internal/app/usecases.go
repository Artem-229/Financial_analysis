package app

import (
	"context"
)

type Usecases struct {
}

func NewUsecases(ctx context.Context, kafka *Kafka) (*Usecases, error) {
	return &Usecases{}, nil
}
