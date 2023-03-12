package presenter

import "net/http"

func JSON(c contextForPresenter, statusCode int, out interface{}) error {
	return c.JSON(statusCode, out)
}

func NoContent(c contextForPresenter) error {
	return c.NoContent(http.StatusNoContent)
}
