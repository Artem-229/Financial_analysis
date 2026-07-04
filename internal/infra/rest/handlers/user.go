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
	LoginUser(ctx context.Context, login string, password string) (string, error)
}

func Registration(usecases Usecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req *dto.RegistrationRequest

		ctx := c.Request.Context()

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := usecases.CreateUser(ctx, convertRegistrationRequestDTOToEntity(req)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "OK",
		})
	}
}

func Login(usecases Usecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req *dto.LoginRequest

		ctx := c.Request.Context()

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := usecases.LoginUser(ctx, req.Login, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":       "OK",
			"access_token": token,
		})
	}
}
