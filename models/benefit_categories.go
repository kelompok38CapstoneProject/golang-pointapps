package models

import (
	"time"

	"gorm.io/gorm"
)

type BenefitCategories struct {
	Id          int    `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
