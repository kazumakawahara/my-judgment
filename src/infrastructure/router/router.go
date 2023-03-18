package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"my-judgment/common/mysql"
)

func Run() {
	e := echo.New()
	db := mysql.Connect()

	// health check用
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	router(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}

func router(e *echo.Echo, db *gorm.DB) {
	RouterForWebGroup(e, db)
}
