package rest

import (
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/rest/handlers"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
	Logger *slog.Logger
}

type ServerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

func NewServer(deps *ServerDeps) *Server {
	router := gin.Default()

	return &Server{
		Engine: router,
		Config: deps.Config,
		Logger: deps.Logger,
	}
}

func (s *Server) Start() error {
	s.Engine.GET("/health", handlers.HealthHandler)

	if err := s.Engine.Run(s.Config.Server.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
