package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/itdevsamurai/gke/simplewebapp/webapp/handler"
)

// Default Server Port
const DEFAULT_SERVER_PORT = ":80"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", handler.SimpleWebHandler{}.Index)

	e.GET("/ping", handler.SimpleWebHandler{}.Ping)

	// Start server
	e.Logger.Fatal(e.Start(DEFAULT_SERVER_PORT))
}
