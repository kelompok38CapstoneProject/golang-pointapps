package route

import (
	"point/constants"
	"point/controllers"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	//users
	//
	e.POST("/singup/", controllers.SingupUserController)
	e.POST("/login/", controllers.LoginUserController)
	e.GET("/user/", controllers.GetAllUserController)
	e.GET("/user/:code", controllers.GetUserControllerCode)
	//jwt_user
	ejwt := e.Group("")
	ejwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	ejwt.PUT("/user/:code", controllers.UpdateUserController)
	ejwt.DELETE("/user/:code", controllers.DeleteUserController)
	return e
}