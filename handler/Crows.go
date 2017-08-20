package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Crows(c echo.Context) error {
	var cry string = "pyokopyoko"
	return c.String(http.StatusOK, cry)
}
