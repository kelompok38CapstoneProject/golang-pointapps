package models

import (
	"time"

	"gorm.io/gorm"
)

type Benefits struct {
	Id                 int                `json:"id"`
	Benefit_Categories Benefit_Categories `json:"benefit_categories"`
	Name               string             `json:"name" `
	Description        string             `json:"description"`
	Price              int                `json:"price"`
	Stock              int                `json:"stock"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
