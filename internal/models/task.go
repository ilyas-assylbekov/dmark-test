package models

import (
	"time"
)

// Task представляет собой модель задачи в системе
type Task struct {
	// ID уникальный идентификатор задачи
	// json: используется для сериализации в API
	// gorm: настройка первичного ключа в базе данных
	ID uint `json:"id" gorm:"primarykey"`

	// Title название или описание задачи
	// json: поле для API
	// gorm: обязательное поле (not null)
	Title string `json:"title" gorm:"not null"`

	// Completed статус выполнения задачи
	// json: поле для API
	// gorm: по умолчанию false
	Completed bool `json:"completed" gorm:"default:false"`

	// CreatedAt время создания задачи
	// Автоматически заполняется GORM
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt время последнего обновления задачи
	// Автоматически обновляется GORM
	UpdatedAt time.Time `json:"updatedAt"`

	// DueDate срок выполнения задачи
	// json: опциональное поле, пропускается если nil
	// Указатель на time.Time позволяет хранить NULL в базе данных
	DueDate *time.Time `json:"dueDate,omitempty"`

	// Priority приоритет задачи
	// json: поле для API
	// gorm: по умолчанию 0
	Priority int `json:"priority" gorm:"default:0"`
}
