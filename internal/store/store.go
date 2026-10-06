package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/example/aurora-orchestrator/internal/agent"
)

// New initializes a Gorm DB and performs migrations.
func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate agent model
	if err := db.AutoMigrate(&agent.Agent{}); err != nil {
		return nil, err
	}

	return db, nil
}
