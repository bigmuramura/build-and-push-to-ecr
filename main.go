package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"build-and-push-ecr/handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.HomePage)
	e.GET("/health", handler.HealthCheck)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
