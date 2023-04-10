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
	r := e.Group("/users")
	r.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	r.GET("", controllers.GetUsersController)
	r.DELETE("/:id", controllers.DeleteUserController)
	r.GET("/:id", controllers.GetUserController)
	r.PUT("/:id", controllers.UpdateUserController)
	b := e.Group("/books")
	b.GET("", controllers.GetBooksController)
	b.POST("", controllers.CreateBookController)
	b.GET("/:id", controllers.GetBookController)
	b.PUT("/:id", controllers.UpdateBookController)
	b.DELETE("/:id", controllers.DeleteBookController)
	e.POST("/login", controllers.LoginUsersController)
	e.POST("/users", controllers.CreateUserController)
	return e
}
