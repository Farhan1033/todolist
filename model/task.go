package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-" binding:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime; autoUpdateTime"`
}
