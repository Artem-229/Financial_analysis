package entities

import "github.com/google/uuid"

type TransactionStatus string

const (
	TransactionStatusPending    TransactionStatus = "pending"
	TransactionStatusComplete   TransactionStatus = "complete"
	TransactionStatusProcessing TransactionStatus = "processing"
)

type Transaction struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Item   string
	Price  int
	Class  string
	Note   string
	Status string
}
