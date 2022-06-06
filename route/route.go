package route

import (
	"point/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	//users
	e.POST("/user/", controller.CreateUserController)
	e.GET("/user/", controller.GetAllUserController)
	e.GET("/user/:code", controller.GetUserControllerCode)
	e.PUT("/user/:code", controller.UpdateUserController)
	e.DELETE("/user/:code", controller.DeleteUserController)
	return e
}
