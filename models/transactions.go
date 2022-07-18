package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	User      Users          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BenefitID int            `json:"benefit_id"`
	Benefit   Benefits       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
