package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

// request POST 'http://127.0.0.1:8080/user/'
func CreateUserController(c echo.Context) error {
	user := models.Users{}
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
