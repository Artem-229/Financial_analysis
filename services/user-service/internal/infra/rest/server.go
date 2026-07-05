package rest

import (
	config2 "financial_assistant/services/user-service/internal/config"
	handlers2 "financial_assistant/services/user-service/internal/infra/rest/handlers"
	"financial_assistant/services/user-service/internal/infra/rest/middleware"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine       *gin.Engine
	Config       *config2.Config
	Creds        *config2.Credentials
	Logger       *slog.Logger
	Usecases     handlers2.UserUsecases
	Transactions handlers2.TransactionsUsecases
}

type ServerDeps struct {
	Config       *config2.Config
	Creds        *config2.Credentials
	Logger       *slog.Logger
	Usecases     handlers2.UserUsecases
	Transactions handlers2.TransactionsUsecases
}

func NewServer(deps *ServerDeps) *Server {
	router := gin.Default()

	return &Server{
		Engine:       router,
		Config:       deps.Config,
		Creds:        deps.Creds,
		Logger:       deps.Logger,
		Usecases:     deps.Usecases,
		Transactions: deps.Transactions,
	}
}

func (s *Server) Start() error {
	s.Engine.GET("/health", handlers2.HealthHandler)
	s.Engine.POST("/registration", handlers2.Registration(s.Usecases))
	s.Engine.POST("/login", handlers2.Login(s.Usecases))

	app := s.Engine.Group("/app", middleware.JWTAuth(s.Creds.JWT.Secret))
	{
		app.POST("/transactions", handlers2.CreateTransaction(s.Transactions))
		app.GET("/transactions", handlers2.ListTransactions(s.Transactions))
		app.GET("/transactions/:id", handlers2.GetTransaction(s.Transactions))
		app.PUT("/transactions/:id", handlers2.UpdateTransaction(s.Transactions))
		app.DELETE("/transactions/:id", handlers2.DeleteTransaction(s.Transactions))
	}

	if err := s.Engine.Run(s.Config.Server.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
