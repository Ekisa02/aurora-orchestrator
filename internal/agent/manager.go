package agent

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

// Manager coordinates agent lifecycle and runs.
type Manager struct {
	db       *gorm.DB
	runs     map[uint]context.CancelFunc
	runsLock sync.Mutex
}

// NewManager creates an agent manager bound to a Gorm DB.
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db:   db,
		runs: make(map[uint]context.CancelFunc),
	}
}

// Create creates and persists a new agent. capabilities may be any JSON-compatible object.
func (m *Manager) Create(a *Agent, caps interface{}) error {
	if a == nil {
		return errors.New("agent is nil")
	}
	if caps != nil {
		b, err := json.Marshal(caps)
		if err != nil {
			return err
		}
		a.Capabilities = b
	}
	return m.db.Create(a).Error
}

// List returns all agents.
func (m *Manager) List() ([]Agent, error) {
	var agents []Agent
	if err := m.db.Find(&agents).Error; err != nil {
		return nil, err
	}
	return agents, nil
}

// Get returns an agent by ID.
func (m *Manager) Get(id uint) (*Agent, error) {
	var a Agent
	if err := m.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

// Execute starts a mock execution for an agent and updates status. This is where real AI orchestration would be integrated.
func (m *Manager) Execute(id uint) error {
	a, err := m.Get(id)
	if err != nil {
		return err
	}

	// If already running, return
	if a.Status == "running" {
		return errors.New("agent already running")
	}

	// Create cancellable context for run control
	ctx, cancel := context.WithCancel(context.Background())

	m.runsLock.Lock()
	m.runs[id] = cancel
	m.runsLock.Unlock()

	// mark running
	if err := m.db.Model(&Agent{}).Where("id = ?", id).Update("status", "running").Error; err != nil {
		return err
	}

	go func() {
		defer func() {
			m.runsLock.Lock()
			delete(m.runs, id)
			m.runsLock.Unlock()
		}()

		// Mock work: multiple steps with checks for cancellation
		steps := 3
		for i := 0; i < steps; i++ {
			select {
			case <-ctx.Done():
				// mark aborted
				_ = m.db.Model(&Agent{}).Where("id = ?", id).Update("status", "aborted").Error
				return
			case <-time.After(500 * time.Millisecond):
				// simulate step
			}
		}

		// Completed
		_ = m.db.Model(&Agent{}).Where("id = ?", id).Update("status", "completed").Error
	}()

	return nil
}

// CancelRun cancels an ongoing run for the agent.
func (m *Manager) CancelRun(id uint) error {
	m.runsLock.Lock()
	cancel, ok := m.runs[id]
	m.runsLock.Unlock()
	if !ok {
		return errors.New("no running job for agent")
	}
	cancel()
	return nil
}
