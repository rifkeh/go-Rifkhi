package route

import (
	"competence/controller"
	m "competence/middleware"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)
func New() *echo.Echo{
	errenv := godotenv.Load()

	if errenv != nil{
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	s := e.Group("")
	s.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	m.LogMiddleware(e)
	s.GET("/items", controller.GetAllBarang)
	s.GET("/item/:id", controller.GetBarang)
	s.POST("/item", controller.CreateBarang)
	s.PUT("/item/:id", controller.UpdateBarang)
	s.DELETE("/item/:id", controller.DeleteBarang)
	s.GET("/items/category/:category_id", controller.GetBarangByKategori)
	e.POST("/register", controller.RegisterAdmin)
	e.POST("/login", controller.LoginAdmin)
	return e
}


