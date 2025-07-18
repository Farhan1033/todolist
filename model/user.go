package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Email      string    `gorm:"unique" json:"email" binding:"required,email"`
	Password   string    `json:"password" binding:"required,min=6"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime; autoUpdateTime"`
	Updated_At time.Time `json:"updated_at" gorm:"autoCreateTime; autoUpdateTime"`
}
