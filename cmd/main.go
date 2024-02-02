package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gormTest/routes"
)

func main() {
	e := echo.New()
	//e.Static("/", "server/static")
	// Middleware for handling CORS, logging, and recovering from panics
	//e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Route for handling image upload
	//	e.POST("/api/images", handleImageUpload)

	// Start the Echo server
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
