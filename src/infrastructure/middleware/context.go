package middleware

import (
	"github.com/labstack/echo/v4"

	"my-judgment/infrastructure/customcontext"
)

func CustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customcontext.CustomContext{Context: c}
			return next(cc)
		}
	}
}
