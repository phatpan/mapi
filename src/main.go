package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Username  string `json:"userName"`
	Password  string `json:"password,omitempty"`
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func getUsers(c echo.Context) error {
	kea := User{
		Firstname: "phatpan",
		Lastname:  "phatpan",
		Username:  "phatpan",
	}
	return c.JSON(http.StatusOK, kea)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index)
	e.GET("/users", getUsers)
	e.Logger.Fatal(e.Start(":1323"))
}
