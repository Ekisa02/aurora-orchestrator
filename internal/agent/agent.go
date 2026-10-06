package agent

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/datatypes"
)

// Agent is the persisted domain model for an autonomous agent.
// Capabilities is stored as JSON for flexibility.
type Agent struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Name         string         `json:"name" gorm:"index;not null"`
	Capabilities datatypes.JSON `json:"capabilities" gorm:"type:json"`
	Status       string         `json:"status" gorm:"type:varchar(32);default:'idle'"`
}
