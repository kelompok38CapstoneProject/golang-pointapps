package route

import (
	constant "point/constants"
	"point/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//users
	e.POST("/singup/", controllers.SingupUserController)
	e.POST("/login/", controllers.LoginUserController)
	e.GET("/users/", controllers.GetAllUserController)
	e.GET("/user/:code", controllers.GetUserControllerCode)
	//jwt_user
	ejwt := e.Group("")
	ejwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	ejwt.PUT("/user/:code", controllers.UpdateUserController)
	ejwt.DELETE("/user/:code", controllers.DeleteUserController)
	//BenefitCategories_controller
	e.POST("/benefitCategorie/", controllers.CreateBenefitCategorieController)
	e.GET("/benefitCategories/", controllers.GetAllBenefitCategorieController)
	e.GET("/benefitCategorie/:code", controllers.GetBenefitCategorieControllerCode)
	e.PUT("/benefitCategorie/:code", controllers.UpdateBenefitCategorieController)
	e.DELETE("/benefitCategorie/:code", controllers.DeleteBenefitCategorieController)
	//benefits_controller
	e.POST("/benefit/", controllers.CreateBenefitsController)
	e.GET("/benefits/", controllers.GetAllBenefitsController)
	e.POST("/benefit/add/:code", controllers.AddBenefitsController)
	e.GET("/benefit/:code", controllers.GetBenefitsControllerCode)
	e.PUT("/benefit/:code", controllers.UpdateBenefitsController)
	e.DELETE("/benefit/:code", controllers.DeleteBenefitsController)
	// transactions
	e.POST("/transaction/", controllers.CreateTransactionsController)
	e.GET("/transactions/", controllers.GetAllTransactionsController)
	e.GET("/transaction/:code", controllers.GetTransactionsControllerCode)
	return e
}
