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
	Note   string
	Price  int
	Class  string
	Status string
}

type SortedTransactions struct {
	Price int      `json:"price"`
	Notes []string `json:"notes"`
	Items []string `json:"items"`
}
