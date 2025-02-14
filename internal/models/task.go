package models

import (
	"time"
)

type Task struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DueDate   *time.Time
	Priority  int `gorm:"default:0"` // (0=normal, 1=high, etc.)
}
