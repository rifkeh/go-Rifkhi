package controller

import (
	"competence/config"
	m "competence/middleware"
	"competence/model"

	"github.com/labstack/echo/v4"
)

func RegisterAdmin(c echo.Context) error{
	admin := model.Admin{}
	c.Bind(&admin)
	if err := config.DB.Save(&admin).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed create admin",
			"error": err.Error(),
	})
	}
	return c.JSON(200, echo.Map{
		"message": "Success create admin",
		"data": admin,
	})
}

func LoginAdmin(c echo.Context) error{
	admin := model.Admin{}
	c.Bind(&admin)
	if err := config.DB.Where("username=? AND password=?", admin.Username, admin.Password).First(&admin).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed get admin",
			"error": err.Error(),
	})
	}
	token, err := m.CreateJWT(admin.ID)
	if err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed create token",
			"error": err.Error(),
	})}
	return c.JSON(200, echo.Map{
		"message": "Login success",
		"token" : token,
	})
}