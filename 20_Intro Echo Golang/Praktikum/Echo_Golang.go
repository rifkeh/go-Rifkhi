package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = []User{
	{
		Id:       1,
		Name:     "Rifkhi",
		Email:    "rifkhiakbar@gmail.com",
		Password: "123",
	},
	{
		Id:       2,
		Name:     "Berliana",
		Email:    "berlianaindah@gmail.com",
		Password: "1234",
	},
	{
		Id:       3,
		Name:     "Alika",
		Email:    "alikazahira@gmail.com",
		Password: "12345",
	},
}

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("Id"))
	if err != nil {
		panic(err)
	}
	var newUsers []User
	for _, u := range users {
		if u.Id == id {
			newUsers = append(newUsers, u)
		}
	}
	return c.JSON(http.StatusBadGateway, map[string]interface{}{
		"messages": "succes get all users with the requested id",
		"users":    newUsers,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("Id"))
	if err != nil {
		panic(err)
	}
	var newUsers []User
	for _, u := range users {
		if u.Id != id {
			newUsers = append(newUsers, u)
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes get all users with the requested id",
		"users":    newUsers,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("Id"))
	nama := c.FormValue("Name")
	email := c.FormValue("Email")
	pass := c.FormValue("Password")
	if err != nil {
		panic(err)
	}
	for i, u := range users {
		if u.Id == id {
			users[i].Name = nama
			users[i].Email = email
			users[i].Password = pass
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes get all users with the requested id",
		"users":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8888"))
}
