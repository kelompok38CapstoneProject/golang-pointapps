package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:efraim@tcp(127.0.0.1:3306)/point?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

}
