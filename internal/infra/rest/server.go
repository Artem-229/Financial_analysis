package rest

import (
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/rest/handlers"
	"financial_assistant/internal/infra/rest/middleware"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine       *gin.Engine
	Config       *config.Config
	Creds        *config.Credentials
	Logger       *slog.Logger
	Usecases     handlers.UserUsecases
	Transactions handlers.TransactionsUsecases
}

type ServerDeps struct {
	Config       *config.Config
	Creds        *config.Credentials
	Logger       *slog.Logger
	Usecases     handlers.UserUsecases
	Transactions handlers.TransactionsUsecases
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
	s.Engine.GET("/health", handlers.HealthHandler)
	s.Engine.POST("/registration", handlers.Registration(s.Usecases))
	s.Engine.POST("/login", handlers.Login(s.Usecases))

	app := s.Engine.Group("/app", middleware.JWTAuth(s.Creds.JWT.Secret))
	{
		app.POST("/transactions", handlers.CreateTransaction(s.Transactions))
		app.GET("/transactions", handlers.ListTransactions(s.Transactions))
		app.GET("/transactions/:id", handlers.GetTransaction(s.Transactions))
		app.PUT("/transactions/:id", handlers.UpdateTransaction(s.Transactions))
		app.DELETE("/transactions/:id", handlers.DeleteTransaction(s.Transactions))
	}

	if err := s.Engine.Run(s.Config.Server.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
