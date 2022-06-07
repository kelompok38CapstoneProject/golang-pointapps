package main

import (
	"point/config"
	route "point/routes"
)

func main() {
	// db, _ := gorm.Open(mysql.Open("gorm.db"), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// })

	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
