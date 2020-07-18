package routes

import (
	"fazt-api-go/middlewares"
	"net/http"

	"fazt-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.Use(middlewares.CORSMiddleware())

	router.GET("/", welcome)

	jwt := router.Group("/")

	jwt.Use(middlewares.Jwt())



	jwt.GET("/users",controllers.GetAllUsers)
		router.POST("/user", controllers.CreateUser)



	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("/todo/:todoId", controllers.GetSingleTodo)
	router.PUT("/todo/:todoId", controllers.EditTodo)
	router.DELETE("/todo/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
