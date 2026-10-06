package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/example/aurora-orchestrator/internal/agent"
)

// Handler groups manager-backed handler functions
type Handler struct {
	mgr *agent.Manager
}

// New constructs handlers using the agent manager.
func New(m *agent.Manager) *Handler {
	return &Handler{mgr: m}
}

// Registered routes live in internal/router, but these functions are public for that wiring.

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) CreateAgent(c *gin.Context) {
	var req struct {
		Name         string      `json:"name" binding:"required"
		Capabilities interface{} `json:"capabilities"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := &agent.Agent{Name: req.Name, Status: "idle"}
	if err := h.mgr.Create(a, req.Capabilities); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, a)
}

func (h *Handler) ListAgents(c *gin.Context) {
	agents, err := h.mgr.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agents)
}

func (h *Handler) GetAgent(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)
	a, err := h.mgr.Get(id)
	if err != nil {
		if err == agent.ErrRecordNotFound || err == nil {
			// Gorm returns gorm.ErrRecordNotFound, but keep generic
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *Handler) ExecuteAgent(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)
	if err := h.mgr.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "started"})
}
