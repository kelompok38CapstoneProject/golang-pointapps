package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"point/config"
	"point/models"

	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/benefits/id'
func GetBenefitsControllerCode(c echo.Context) error {
	benefitsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefits models.Benefits
	if err := config.DB.First(&benefits, benefitsId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if benefits.ID == 0 {
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if err := config.DB.Preload("BenefitCategory").Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}

// request GET 'http://127.0.0.1:8080/benefits/'
func GetAllBenefitsController(c echo.Context) error {
	var benefits []models.Benefits
	if err := config.DB.Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("BenefitCategory").Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}

// request PUT 'http://127.0.0.1:8080/benefits/add/id'
func AddBenefitsController(c echo.Context) error {
	benefitsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefits models.Benefits
	var reqBenefits models.Benefits
	if err := config.DB.First(&benefits, benefitsId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if benefits.ID == 0 {
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if err := c.Bind(&reqBenefits); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	benefits.Stock += reqBenefits.Stock
	if err := config.DB.Save(&benefits).Error; err != nil {
		fmt.Println(err)

		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("BenefitCategory").Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}

// request POST 'http://127.0.0.1:8080/benefits/'
func CreateBenefitsController(c echo.Context) error {
	benefits := models.Benefits{}
	if err := c.Bind(&benefits); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("BenefitCategory").Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}

// request PUT 'http://127.0.0.1:8080/benefits/id'
func UpdateBenefitsController(c echo.Context) error {
	benefitsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefits models.Benefits
	if err := config.DB.First(&benefits, benefitsId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if benefits.ID == 0 {
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if err := c.Bind(&benefits); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Preload("BenefitCategory").Find(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}

// request DELETE 'http://127.0.0.1:8080/benefits/id'
func DeleteBenefitsController(c echo.Context) error {
	benefitsId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefits models.Benefits
	if err := config.DB.First(&benefits, benefitsId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if benefits.ID == 0 {
		return c.String(http.StatusNotFound, "benefits not found")
	}
	if err := config.DB.Delete(&benefits).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefits)
}
