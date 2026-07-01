package rest

import (
	"financial_assistant/internal/infra/rest/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()

	router.GET("/health", handlers.HealthHandler)
	return &Server{
		Engine: router,
	}
}
