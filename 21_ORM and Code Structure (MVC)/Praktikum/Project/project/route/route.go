package route

import (
	"project/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.GET("/users/:id", controllers.GetUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.GET("/books", controllers.GetBooksController)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books/:id", controllers.GetBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.GET("/blogs", controllers.GetBlogsController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	return e
}
