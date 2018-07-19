package controller

import (
	"github.com/labstack/echo"
)

type baseController struct {
}

func (b baseController) throwError(code int, message string) error {
	return echo.NewHTTPError(code, message)
}
