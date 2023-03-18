package presenter

import "net/http"

type contextForPresenter interface {
	JSON(int, interface{}) error
	NoContent(code int) error
	Request() *http.Request
}
