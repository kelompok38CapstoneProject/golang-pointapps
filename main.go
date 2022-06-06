package main

import (
	"point/route"
	"point/config"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}