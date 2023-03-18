package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"my-judgment/common/config"
)

func DBMiddleware(conn *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			ctx := context.WithValue(r.Context(), config.DBKey, conn)
			*r = *r.WithContext(ctx)

			return next(c)
		}
	}
}
