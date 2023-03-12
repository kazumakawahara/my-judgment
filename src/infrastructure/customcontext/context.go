package customcontext

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
}

func (c CustomContext) SessionID() string {
	return c.Get("sessionID").(string)
}

func (c CustomContext) UserID() int {
	return c.Get("userID").(int)
}
