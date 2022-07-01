package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"point/config"
	"point/models"

	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/transactions/id'
func GetTransactionsControllerCode(c echo.Context) error {
	transactionsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var transactions models.Transactions
	if err := config.DB.First(&transactions, transactionsId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "transactions not found")
	}
	if transactions.ID == 0 {
		return c.String(http.StatusNotFound, "transactions not found")
	}
	if err := config.DB.Preload("User").Preload("Benefit.BenefitCategory").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, transactions)
}

// request GET 'http://127.0.0.1:8080/transactions/id/users'
func GetTransactionsUserControllerCode(c echo.Context) error {
	transactionsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var transactions []models.Transactions
	if err := config.DB.Where("user_id=?", transactionsId).Preload("User").Preload("Benefit.BenefitCategory").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if len(transactions) == 0 {
		return c.String(http.StatusNotFound, "transactions not found")
	}
	return c.JSON(http.StatusOK, transactions)
}

// request GET 'http://127.0.0.1:8080/transactions/'
func GetAllTransactionsController(c echo.Context) error {
	var transactions []models.Transactions
	if err := config.DB.Preload("User").Preload("Benefit.BenefitCategory").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, transactions)
}

// request POST 'http://127.0.0.1:8080/transaction/'
func CreateTransactionsController(c echo.Context) error {
	transactions := models.Transactions{}

	if err := c.Bind(&transactions); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 1")
	}
	// Get User By ID
	var user models.Users
	if err := config.DB.First(&user, transactions.UserID).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 2")
	}
	// Get Benefit By ID
	var benefit models.Benefits
	if err := config.DB.First(&benefit, transactions.BenefitID).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 2")
	}
	// Reduce Point User
	if err := config.DB.Model(&user).Where("id = ?", user.ID).Update("point", user.Point-benefit.Price).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 3")
	}
	// Point User < Benefit.Price
	if user.Point < benefit.Price {
		return c.String(http.StatusBadRequest, "Not Enough Point")
	}
	// Stock = 0
	if benefit.Stock == 0 {
			return c.String(http.StatusInternalServerError, "Stock Out")
	} else {
		// Reduce Benefit Stock By 1
		if err := config.DB.Model(&benefit).Update("stock", benefit.Stock-1).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error 4")
		}
	}
	if err := config.DB.Create(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Preload("User").Preload("Benefit.BenefitCategory").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 5")
	}

	return c.JSON(http.StatusOK, transactions)
}
