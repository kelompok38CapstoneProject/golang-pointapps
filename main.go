package main

import (
	"point/config"
	_middlewares "point/middlewares"
	route "point/routes"
)

func main() {

	config.InitDB()
	e := route.New()
	_middlewares.Log(e)
	e.Logger.Fatal(e.Start(":8080"))
}
