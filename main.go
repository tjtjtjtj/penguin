package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tjtjtjtj/peguin/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/penguin/v1/Crows", handler.Crows)

	e.Logger.Fatal(e.Start(":1323"))
}
