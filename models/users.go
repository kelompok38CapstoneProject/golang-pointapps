package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        int            `json:"id"`
	Name      string         `json:"name" `
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Role      string         `json:"role"`
	Point     uint           `json:"point"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UsersResponse struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama" `
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
