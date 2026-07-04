package handlers

import (
	"context"
	"financial_assistant/internal/entities"
	"financial_assistant/internal/infra/rest/handlers/dto"
	"financial_assistant/internal/infra/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionsUsecases interface {
	UpsertTransaction(ctx context.Context, transaction *entities.Transaction) error
	GetTransaction(ctx context.Context, id string, userID string) (*entities.Transaction, error)
	ListTransactions(ctx context.Context, userID string) ([]entities.Transaction, error)
	DeleteTransaction(ctx context.Context, id string, userID string) error
}

func CreateTransaction(usecases TransactionsUsecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dto.CreateTransactionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, err := uuid.Parse(c.GetString(middleware.UserIDKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
			return
		}

		transaction := convertTransactionDTOToEntity(&req, userID)
		if err := usecases.UpsertTransaction(c.Request.Context(), transaction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, convertTransactionEntityToDTO(transaction))
	}
}

func UpdateTransaction(usecases TransactionsUsecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
			return
		}

		var req dto.UpdateTransactionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, err := uuid.Parse(c.GetString(middleware.UserIDKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
			return
		}

		transaction := convertUpdateTransactionDTOToEntity(&req, id, userID)
		if err := usecases.UpsertTransaction(c.Request.Context(), transaction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, convertTransactionEntityToDTO(transaction))
	}
}

func GetTransaction(usecases TransactionsUsecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := c.GetString(middleware.UserIDKey)

		transaction, err := usecases.GetTransaction(c.Request.Context(), id, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, convertTransactionEntityToDTO(transaction))
	}
}

func ListTransactions(usecases TransactionsUsecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		userID := c.GetString(middleware.UserIDKey)

		transactions, err := usecases.ListTransactions(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := make([]*dto.TransactionResponse, 0, len(transactions))
		for i := range transactions {
			response = append(response, convertTransactionEntityToDTO(&transactions[i]))
		}

		c.JSON(http.StatusOK, response)
	}
}

func DeleteTransaction(usecases TransactionsUsecases) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := c.GetString(middleware.UserIDKey)

		if err := usecases.DeleteTransaction(c.Request.Context(), id, userID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
	}
}
