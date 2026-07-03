package handlers

import (
	"context"
	"financial_assistant/internal/entities"
	"financial_assistant/internal/infra/rest/handlers/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usecases interface {
	CreateUser(ctx context.Context, user *entities.User) error
}

func Registration(usecases Usecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req *dto.User

		ctx := c.Request.Context()

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := usecases.CreateUser(ctx, convertDTOToEntity(req)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, gin.H{"user": convertDTOToEntity(req)})
	}
}
