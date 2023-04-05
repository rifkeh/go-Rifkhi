package controllers

import (
	"net/http"
	"project/config"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting all blog",
		"blogs":   blogs,
	})
}

func CreateBlogController(c echo.Context) error {
	var blogs models.Blog
	c.Bind(&blogs)

	if err := config.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"blogs":   blogs,
	})
}

func DeleteBlogController(c echo.Context) error {
	var blogs models.Blog
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Delete(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succrss delete blog",
	})
}

func GetBlogController(c echo.Context) error {
	var blogs models.Blog
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes get specified user",
		"blog":    blogs,
	})
}

func UpdateBlogController(c echo.Context) error {
	var blogs models.Blog
	id := c.Param("id")
	c.Bind(&blogs)
	if err := config.DB.Where("id = ?", id).Updates(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
	})
}
