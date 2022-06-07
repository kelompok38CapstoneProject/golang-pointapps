package models

import (
	"time"

	"gorm.io/gorm"
)

type Benefit_Categories struct {
	Id          int    `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
