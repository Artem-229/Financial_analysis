package handlers

import (
	"financial_assistant/internal/entities"
	"financial_assistant/internal/infra/rest/handlers/dto"

	"github.com/google/uuid"
)

func convertRegistrationRequestDTOToEntity(req *dto.RegistrationRequest) *entities.User {
	rawID, err := uuid.NewUUID()
	if err != nil {
		return &entities.User{}
	}

	return &entities.User{
		Id:       rawID,
		Username: req.Username,
		Password: req.Password,
		Login:    req.Login,
	}
}

func convertTransactionDTOToEntity(req *dto.CreateTransactionRequest, userID uuid.UUID) *entities.Transaction {
	rawID, err := uuid.NewUUID()
	if err != nil {
		return &entities.Transaction{}
	}

	return &entities.Transaction{
		ID:     rawID,
		UserID: userID,
		Price:  req.Price,
		Item:   req.Item,
		Class:  req.Class,
	}
}

func convertUpdateTransactionDTOToEntity(req *dto.UpdateTransactionRequest, id uuid.UUID, userID uuid.UUID) *entities.Transaction {
	return &entities.Transaction{
		ID:     id,
		UserID: userID,
		Price:  req.Price,
		Item:   req.Item,
		Class:  req.Class,
	}
}

func convertTransactionEntityToDTO(transaction *entities.Transaction) *dto.TransactionResponse {
	return &dto.TransactionResponse{
		ID:    transaction.ID.String(),
		Item:  transaction.Item,
		Price: transaction.Price,
		Class: transaction.Class,
	}
}
