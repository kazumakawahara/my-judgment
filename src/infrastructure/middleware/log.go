package middleware

import (
	"log"

	"github.com/labstack/echo/v4"

	"my-judgment/common/config"
	mjerr2 "my-judgment/common/mjerr"
)

func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				log.Printf("%+v", err)
				return err
			}

			ctx := c.Request().Context()
			err, ok := ctx.Value(config.LogKey).(error)
			if !ok {
				// 処理成功
				return nil
			}

			apoErr := mjerr2.AsApoError(err)
			if apoErr == nil {
				log.Printf("failed in type assertion for apoError: %+v", err)
				return nil
			}

			var logLevel mjerr2.Level
			if appErr := apoErr.OriginError(); appErr != nil {
				logLevel = appErr.Level()
			} else {
				// デフォルトエラーレベル
				logLevel = mjerr2.LevelFatal
				log.Print("failed in type assertion for OriginError")
			}

			log.Printf("[level]: %s\n    %+v", logLevel, apoErr)

			return nil
		}
	}
}
