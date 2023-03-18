package middleware

import (
	"github.com/labstack/echo/v4"

	"my-judgment/common/apperr"
	mjerr2 "my-judgment/common/mjerr"
	"my-judgment/interfaces/presenter"
)

func VerifyApiVersion() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			version := c.Request().Header.Get("my-judgment-api-version")

			// Api Version検証
			if value, exist := apiVersion[version]; value && exist {
				return next(c)
			}

			return presenter.ErrorJSON(
				c,
				mjerr2.Wrap(
					nil,
					mjerr2.WithOriginError(apperr.Gone),
					mjerr2.WithLogDetail(
						map[string]interface{}{
							"version": version,
						},
					),
				),
			)
		}
	}
}

var apiVersion map[string]bool = map[string]bool{
	"1.0": true,
}
