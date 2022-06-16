package config

import (
	"fmt"
	"point/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db := "pointapps"
	connectionString := fmt.Sprintf("root:qqwerty@tcp(127.0.0.1:3306)/%s?parseTime=True", db)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Users{}, &models.BenefitCategories{}, &models.Benefits{}, &models.Transactions{})
}
