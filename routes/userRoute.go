package routes

import (
	"echi/controllers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/users", controllers.AllUser)
	e.POST("/user", controllers.CreateUser)
}
