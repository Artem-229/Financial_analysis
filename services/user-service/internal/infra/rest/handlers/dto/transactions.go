package dto

type CreateTransactionRequest struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
	Note  string `json:"note"`
}

type UpdateTransactionRequest struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
	Note  string `json:"note"`
}

type TransactionResponse struct {
	ID    string `json:"id"`
	Item  string `json:"item"`
	Price int    `json:"price"`
	Class string `json:"class"`
	Note  string `json:"note"`
}
