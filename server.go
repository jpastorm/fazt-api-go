package main

import (
	"echi/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.Routes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
