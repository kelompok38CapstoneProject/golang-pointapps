package models

import (
	"time"
	// "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Benefits struct {
	Id                int               `json:"id"`
	BenefitCategoryId int               `json:"benefit_category_id"`
	BenefitCategory   BenefitCategories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name              string            `json:"name" `
	Description       string            `json:"description"`
	Price             uint              `json:"price"`
	Stock             uint              `json:"stock"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
