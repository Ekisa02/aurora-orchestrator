package router

import (
	"github.com/gin-gonic/gin"
	"github.com/example/aurora-orchestrator/internal/handlers"
)

// New wires up routes and middleware
func New(h *handlers.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", h.Health)

	agents := r.Group("/agents")
	{
		agents.POST("", h.CreateAgent)
		agents.GET("", h.ListAgents)
		agents.GET(":id", h.GetAgent)
		agents.POST(":id/execute", h.ExecuteAgent)
	}

	return r
}
