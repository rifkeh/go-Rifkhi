package controllers

import (
	"net/http"
	"project/config"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting all book",
		"books":   books,
	})
}

func CreateBookController(c echo.Context) error {
	var books models.Book
	c.Bind(&books)

	if err := config.DB.Save(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	var books []models.Book

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get specified user",
		"books":   books,
	})
}

func DeleteBookController(c echo.Context) error {
	var books models.Book

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).Delete(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateBookController(c echo.Context) error {
	var books models.Book
	c.Bind(&books)
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Updates(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
