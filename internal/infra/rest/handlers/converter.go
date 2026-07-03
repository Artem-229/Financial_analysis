package handlers

import (
	"financial_assistant/internal/entities"
	"financial_assistant/internal/infra/rest/handlers/dto"

	"github.com/google/uuid"
)

func convertDTOToEntity(req *dto.User) *entities.User {
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
