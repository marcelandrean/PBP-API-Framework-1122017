package main

import (
	"net/http"

	"echo-rest/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	controllers.Init()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":8888"))
}
