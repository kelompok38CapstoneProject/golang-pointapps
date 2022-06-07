package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id        int      `json:"id"`
	Users     Users    `json:"users"`
	Benefits  Benefits `json:"benefits"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
