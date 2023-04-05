package route

import (
	"project/constants"
	"project/controllers"
	m "project/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)
	r := e.Group("/jwt")
	r.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	e.POST("/login", controllers.LoginUsersController)
	r.GET("/users", controllers.GetUsersController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.GET("/users/:id", controllers.GetUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.GET("/books", controllers.GetBooksController)
	r.POST("/books", controllers.CreateBookController)
	r.GET("/books/:id", controllers.GetBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	e.POST("/users", controllers.CreateUserController)
	return e
}
