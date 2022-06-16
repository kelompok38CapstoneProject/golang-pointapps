package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	Id        int      `json:"id"`
	UserId    int      `json:"user_id"`
	User      Users    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BenefitId int      `json:"benefit_id"`
	Benefit   Benefits `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
