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
	if err := config.DB.Preload("Users").Preload("Benefits").Find(&transactions).Error; err != nil {
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
	if err := config.DB.Preload("Users").Preload("Benefits").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, transactions)
}

// request POST 'http://127.0.0.1:8080/transactions/'
func CreateTransactionsController(c echo.Context) error {
	transactions := models.Transactions{}
	if err := c.Bind(&transactions); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("Users").Preload("Benefits").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, transactions)
}
