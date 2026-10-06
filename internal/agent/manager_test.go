package agent

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestManager(t *testing.T) *Manager {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.AutoMigrate(&Agent{}); err != nil {
		t.Fatalf("migrate failed: %v", err)
	}
	return NewManager(db)
}

func TestCreateAndExecute(t *testing.T) {
	m := newTestManager(t)
	agent := &Agent{Name: "test-agent"}
	if err := m.Create(agent, map[string]bool{"can_search": true}); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	// ensure initial status
	a, err := m.Get(agent.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if a.Status != "idle" {
		t.Fatalf("expected idle status, got %s", a.Status)
	}

	if err := m.Execute(agent.ID); err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	// Wait for completion (mock sleeps for ~1.5s across steps)
	waited := 0
	for {
		a, err := m.Get(agent.ID)
		if err != nil {
			t.Fatalf("get failed while waiting: %v", err)
		}
		if a.Status == "completed" {
			break
		}
		if waited > 10 {
			t.Fatalf("timed out waiting for completion, last status=%s", a.Status)
		}
		waited++
		time.Sleep(200 * time.Millisecond)
	}
}
