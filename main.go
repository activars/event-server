package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Post("/event", func(c echo.Context) error {
		//
		c.Logger().Info("message received, sending to pubsub");
		return c.JSON(http.StatusAccepted, make(map[string]string))
	})
	app.Run(standard.New(":3000"))
}