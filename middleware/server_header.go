package middleware

import (
	"github.com/labstack/echo"
)

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "ECHO/3.0")
			return next(c)
		}
	}
}
