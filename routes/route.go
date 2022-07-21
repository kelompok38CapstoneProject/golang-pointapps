package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	constant "point/constants"
	"point/controllers"
	"point/middlewares"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	ejwt := e.Group("")
	//users
	ejwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	e.POST("/singup/", controllers.SingupUserController)
	e.POST("/login/user", controllers.LoginUserController)
	e.POST("/login/admin", controllers.LoginUserController)

	ejwt.POST("/users/role/admin/", controllers.CreateUserRoleAdminController, middlewares.AdminRole)
	ejwt.POST("/users/role/user/", controllers.CreateUserRoleUserController, middlewares.AdminRole)
	ejwt.GET("/users/role/admin/", controllers.GetAllUserRoleAdminController, middlewares.AdminRole)
	ejwt.GET("/users/role/user/", controllers.GetAllUserRoleUserController)
	ejwt.GET("/users/:code", controllers.GetUserControllerCode)
	ejwt.PUT("/addpointusers/:code", controllers.AddPointUserController, middlewares.AdminRole)
	ejwt.PUT("/users/:code", controllers.UpdateUserController)
	ejwt.DELETE("/users/:code", controllers.DeleteUserController, middlewares.AdminRole)
	ejwt.GET("/users/", controllers.GetAllUserController, middlewares.AdminRole)
	//BenefitCategories_controller
	ejwt.POST("/benefitCategories/", controllers.CreateBenefitCategorieController, middlewares.AdminRole)
	ejwt.GET("/benefitCategories/", controllers.GetAllBenefitCategorieController)
	ejwt.GET("/benefitCategories/:code", controllers.GetBenefitCategorieControllerCode)
	ejwt.PUT("/benefitCategories/:code", controllers.UpdateBenefitCategorieController, middlewares.AdminRole)
	ejwt.DELETE("/benefitCategorie/:code", controllers.DeleteBenefitCategorieController, middlewares.AdminRole)
	//benefits_controller
	ejwt.POST("/benefits/", controllers.CreateBenefitsController, middlewares.AdminRole)
	ejwt.GET("/benefits/", controllers.GetAllBenefitsController)
	ejwt.GET("/benefits/benefitCategoryId/:code", controllers.GetBenefitCategoryIDControllerCode)
	ejwt.POST("/benefits/add/:code", controllers.AddBenefitsController, middlewares.AdminRole)
	ejwt.GET("/benefits/:code", controllers.GetBenefitsControllerCode)
	ejwt.PUT("/benefits/:code", controllers.UpdateBenefitsController, middlewares.AdminRole)
	ejwt.DELETE("/benefits/:code", controllers.DeleteBenefitsController, middlewares.AdminRole)
	// transactions
	ejwt.POST("/transactions/", controllers.CreateTransactionsController)
	ejwt.GET("/transactions/", controllers.GetAllTransactionsController)
	ejwt.GET("/transactions/:code", controllers.GetTransactionsControllerCode)
	ejwt.GET("/transactions/:code/users", controllers.GetTransactionsUserControllerCode)
	return e
}
