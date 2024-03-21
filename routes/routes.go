package routes

import (
	"net/http"

	"echo-rest/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	return e
}
