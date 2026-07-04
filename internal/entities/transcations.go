package entities

import "github.com/google/uuid"

type Transaction struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Item   string
	Price  int
	Class  string
}
