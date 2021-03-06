package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"point/config"
	"point/models"

	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/benefitCategorie/id'
func GetBenefitCategorieControllerCode(c echo.Context) error {
	benefitCategorieId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefitCategorie models.BenefitCategories
	if err := config.DB.First(&benefitCategorie, benefitCategorieId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	if benefitCategorie.ID == 0 {
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	return c.JSON(http.StatusOK, benefitCategorie)
}

// request GET 'http://127.0.0.1:8080/benefitCategorie/'
func GetAllBenefitCategorieController(c echo.Context) error {
	var benefitCategorie []models.BenefitCategories
	if err := config.DB.Find(&benefitCategorie).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	// return c.JSON(http.StatusOK, benefitCategorie)
	return c.JSON(200, map[string]interface{}{"benefitCategories": benefitCategorie})
}

// request POST 'http://127.0.0.1:8080/benefitCategorie/'
func CreateBenefitCategorieController(c echo.Context) error {
	benefitCategorie := models.BenefitCategories{}
	if err := c.Bind(&benefitCategorie); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&benefitCategorie).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefitCategorie)
}

// request PUT 'http://127.0.0.1:8080/benefitCategorie/id'
func UpdateBenefitCategorieController(c echo.Context) error {
	benefitCategorieId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefitCategorie models.BenefitCategories
	if err := config.DB.First(&benefitCategorie, benefitCategorieId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	if benefitCategorie.ID == 0 {
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	if err := c.Bind(&benefitCategorie); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Save(&benefitCategorie).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefitCategorie)
}

// request DELETE 'http://127.0.0.1:8080/benefitCategorie/id'
func DeleteBenefitCategorieController(c echo.Context) error {
	benefitCategorieId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var benefitCategorie models.BenefitCategories
	if err := config.DB.First(&benefitCategorie, benefitCategorieId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	if benefitCategorie.ID == 0 {
		return c.String(http.StatusNotFound, "benefitCategorie not found")
	}
	if err := config.DB.Delete(&benefitCategorie).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, benefitCategorie)
}
