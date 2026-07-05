package dto

type CreateTransactionRequest struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
}

type UpdateTransactionRequest struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
}

type TransactionResponse struct {
	ID    string `json:"id"`
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
}
