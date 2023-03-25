package mockcontext

import "github.com/labstack/echo/v4"

type MockContext struct {
	echo.Context
}

func (c *MockContext) UserID() int {
	return 0
}
