package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Point     int    `json:"point"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UsersResponse struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama" `
	Email string `json:"email"`
	Token    string `json:"token"`
}