package handler

import (
	"net/http"
)

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	NoContent(code int) error
	Param(string) string
	QueryParam(string) string
	String(int, string) error
	FormValue(string) string
	Request() *http.Request
	Set(key string, val interface{})
	SetCookie(cookie *http.Cookie)
	UserID() int
}
