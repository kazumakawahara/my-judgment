package middleware

import "github.com/labstack/echo/v4"

func VerifyToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// サンプルAPIなので値は決めうち
			c.Set("userID", 999)
			return next(c)
		}
	}
}
