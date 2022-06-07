package route

import (
	"point/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	//users
	e.POST("/user/", controllers.CreateUserController)
	e.GET("/user/", controllers.GetAllUserController)
	e.GET("/user/:code", controllers.GetUserControllerCode)
	e.PUT("/user/:code", controllers.UpdateUserController)
	e.DELETE("/user/:code", controllers.DeleteUserController)
	return e
}
