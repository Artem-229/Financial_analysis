package app

import "context"

type Usecases struct {
}

func NewUsecases(ctx context.Context) *Usecases {
	return &Usecases{}
}
