package main

import (
	"net/http"

	"github.com/oxlb/GoLangPostgresHelloWorld/controller"
	"github.com/oxlb/GoLangPostgresHelloWorld/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Routes
	e.GET("/", hello)
	e.GET("/student", controller.GetStudents)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
