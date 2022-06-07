package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"point/middleware"
	"point/config"
	"point/models"
	
	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/user/id'
func GetUserControllerCode(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user models.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "user not found")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

// request GET 'http://127.0.0.1:8080/user/'
func GetAllUserController(c echo.Context) error {
	var user []models.Users
	if err := config.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request POST 'http://127.0.0.1:8080/singup/'
func SingupUserController(c echo.Context) error {
	user := models.Users{}
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal singup")
	}
	userResponse := models.UsersResponse{user.Id, user.Name, user.Email, token}
	return c.JSON(http.StatusOK, userResponse)
}

// request PUT 'http://127.0.0.1:8080/user/id'
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user models.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "user not found")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request DELETE 'http://127.0.0.1:8080/user/id'
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user models.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "user not found")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request POST 'http://127.0.0.1:8080/login/'
func LoginUserController(c echo.Context) error {
	user := models.Users{}
	// fmt.Printf("user sebelum bind %#v\n", user)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("user setelah bind %#v\n", user)
	fmt.Printf("Before insert: %#v\n", user)
	if err := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login")
	}

	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login")
	}
	userResponse := models.UsersResponse{user.Id, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, userResponse)
}
