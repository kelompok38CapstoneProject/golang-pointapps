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
	if transactions.Id == 0 {
		return c.String(http.StatusNotFound, "transactions not found")
	}
	if err := config.DB.Preload("User").Preload("Benefit").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, transactions)
}

// request GET 'http://127.0.0.1:8080/transactions/'
func GetAllTransactionsController(c echo.Context) error {
	var transactions []models.Transactions
	if err := config.DB.Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("User").Preload("Benefit").Find(&transactions).Error; err != nil {
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
	// Get Benefit By ID
	var benefit models.Benefits
	if err := config.DB.First(&benefit, transactions.BenefitId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 2")
	}
	// Reduce Benefit Stock By 1
	if err := config.DB.Model(&benefit).Update("stock", benefit.Stock-1).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 3")
	}
	if err := config.DB.Preload("User").Preload("Benefit").Preload("BenefitCategory").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error 3")
	}
	return c.JSON(http.StatusOK, transactions)
}
