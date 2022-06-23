package config

import (
	"point/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	connectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "root:efraim@tcp(127.0.0.1:3306)/point?charset=utf8&parseTime=True&loc=Local"
	}
	

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
