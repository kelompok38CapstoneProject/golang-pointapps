package controllers

import (
	"fmt"
	"net/http"
	"point/config"
	"point/middlewares"
	"point/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

// request GET 'http://127.0.0.1:8080/users/'
func GetAllUserController(c echo.Context) error {
	var user []models.Users
	if err := config.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, user)
}

// request POST 'http://127.0.0.1:8080/singup/'
func SingupUserController(c echo.Context) error {
	var reqUser models.Users

	if err := c.Bind(&reqUser); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	var email string
	if err := config.DB.Table("users").Select("email").Where("email=?", reqUser.Email).Find(&email).Error; err != nil {
		return err
	}
	if email != "" {
		return c.String(http.StatusInternalServerError, "email used")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
	reqUser.Password = string(hashedPassword)
	if err := config.DB.Save(&reqUser).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	token, err := middlewares.CreateToken(reqUser.ID, reqUser.Name)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal singup")
	}
	userResponse := models.UsersResponse{reqUser.ID, reqUser.Name, reqUser.Email, token}
	return c.JSON(http.StatusOK, userResponse)
}

// request POST 'http://127.0.0.1:8080/login/'
func LoginUserController(c echo.Context) error {
	user := models.Users{}
	requser := models.Users{}
	// fmt.Printf("user sebelum bind %#v\n", user)
	if err := c.Bind(&requser); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// fmt.Printf("user setelah bind %#v\n", user)
	fmt.Printf("Before insert: %#v\n", user)

	if err := config.DB.Where("email=?", requser.Email).First(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requser.Password))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login1")
	}
	token, err := middlewares.CreateToken(user.ID, user.Name)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login2")
	}
	userResponse := models.UsersResponse{user.ID, user.Name, user.Email, token}

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
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, user)
}

// request PUT 'http://127.0.0.1:8080/user/id'
func AddPointUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user models.Users
	var reqUser models.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "user not found")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := c.Bind(&reqUser); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	user.Point = user.Point + reqUser.Point
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
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
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, user)
}
