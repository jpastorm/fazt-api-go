package routes

import (
	"echi/controllers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {

	g := e.Group("/user")
	g.GET("", controllers.AllUser)
	g.POST("", controllers.CreateUser)
	g.POST("/login", controllers.LoginUser)
	g.GET("/search/:value", controllers.SearchUser)

}
