package server

import (
	user "angmorning.com/internal/services/users/presentation"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	healthCheckHandler *HealthCheckHandler,
	userHandler *user.UserHandler,
) *Server {
	engine := gin.Default()

	// routing
	engine.GET("/health", healthCheckHandler.check)

	userGroup := engine.Group("/users")
	userHandler.Router(userGroup)

	return &Server{
		engine: engine,
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}