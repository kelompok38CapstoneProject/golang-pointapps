package models

import (
	"time"
	// "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Benefits struct {
	ID                int               `json:"id"`
	BenefitCategoryID int               `json:"benefit_category_id"`
	BenefitCategory   BenefitCategories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name              string            `json:"name" `
	Description       string            `json:"description"`
	Price             uint              `json:"price"`
	Stock             uint              `json:"stock"`
	CreatedAt         time.Time         `json:"-"`
	UpdatedAt         time.Time         `json:"-"`
	DeletedAt         gorm.DeletedAt    `gorm:"index"`
}
