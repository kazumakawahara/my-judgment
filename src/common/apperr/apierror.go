package apperr

import (
	"errors"
	"net/http"

	mjerr2 "my-judgment/common/mjerr"
)

type apiError struct {
	code       string
	statusCode int
	detail     Detail
	level      mjerr2.Level
}

type Detail interface {
	IsDetail() bool
}

func (e *apiError) Error() string {
	return e.code
}

func (e *apiError) StatusCode() int {
	return e.statusCode
}

func (e *apiError) Detail() Detail {
	return e.detail
}

func (e *apiError) WithDetail(detail Detail) *apiError {
	return &apiError{
		code:       e.code,
		statusCode: e.statusCode,
		detail:     detail,
		level:      e.level,
	}
}

func (e *apiError) Level() mjerr2.Level {
	return e.level
}

func AsAPIError(apoErr mjerr2.OriginError) *apiError {
	var e *apiError
	if errors.As(apoErr, &e) {
		return e
	}

	return nil
}

func newBadRequest(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusBadRequest,
		level:      mjerr2.LevelError,
	}
}

func newUnauthorized(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusUnauthorized,
		level:      mjerr2.LevelWarning,
	}
}

func newNotFound(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusNotFound,
		level:      mjerr2.LevelWarning,
	}
}

func newConflict(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusConflict,
		level:      mjerr2.LevelWarning,
	}
}

func newGone(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusGone,
		level:      mjerr2.LevelError,
	}
}

func newInternalServerError(code string) *apiError {
	return &apiError{
		code:       code,
		statusCode: http.StatusInternalServerError,
		level:      mjerr2.LevelFatal,
	}
}
