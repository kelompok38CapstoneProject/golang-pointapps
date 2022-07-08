package models

import (
	"time"

	"gorm.io/gorm"
)

type BenefitCategories struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" `
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
