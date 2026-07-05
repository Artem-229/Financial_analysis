package handlers

import (
	entities2 "financial_assistant/services/user-service/internal/entities"
	dto2 "financial_assistant/services/user-service/internal/infra/rest/handlers/dto"

	"github.com/google/uuid"
)

func convertRegistrationRequestDTOToEntity(req *dto2.RegistrationRequest) *entities2.User {
	rawID, err := uuid.NewUUID()
	if err != nil {
		return &entities2.User{}
	}

	return &entities2.User{
		Id:       rawID,
		Username: req.Username,
		Password: req.Password,
		Login:    req.Login,
	}
}

func convertTransactionDTOToEntity(req *dto2.CreateTransactionRequest, userID uuid.UUID) *entities2.Transaction {
	rawID, err := uuid.NewUUID()
	if err != nil {
		return &entities2.Transaction{}
	}

	return &entities2.Transaction{
		ID:     rawID,
		UserID: userID,
		Price:  req.Price,
		Item:   req.Item,
		Class:  req.Class,
	}
}

func convertUpdateTransactionDTOToEntity(req *dto2.UpdateTransactionRequest, id uuid.UUID, userID uuid.UUID) *entities2.Transaction {
	return &entities2.Transaction{
		ID:     id,
		UserID: userID,
		Price:  req.Price,
		Item:   req.Item,
		Class:  req.Class,
	}
}

func convertTransactionEntityToDTO(transaction *entities2.Transaction) *dto2.TransactionResponse {
	return &dto2.TransactionResponse{
		ID:    transaction.ID.String(),
		Item:  transaction.Item,
		Price: transaction.Price,
		Class: transaction.Class,
	}
}
