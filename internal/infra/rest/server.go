package rest

import (
	"financial_assistant/internal/config"
	"financial_assistant/internal/infra/rest/handlers"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine   *gin.Engine
	Config   *config.Config
	Logger   *slog.Logger
	Usecases handlers.Usecases
}

type ServerDeps struct {
	Config   *config.Config
	Logger   *slog.Logger
	Usecases handlers.Usecases
}

func NewServer(deps *ServerDeps) *Server {
	router := gin.Default()

	return &Server{
		Engine:   router,
		Config:   deps.Config,
		Logger:   deps.Logger,
		Usecases: deps.Usecases,
	}
}

func (s *Server) Start() error {

	s.Engine.GET("/health", handlers.HealthHandler)
	s.Engine.POST("/registration", handlers.Registration(s.Usecases))
	s.Engine.POST("/login", handlers.Login(s.Usecases))

	if err := s.Engine.Run(s.Config.Server.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
