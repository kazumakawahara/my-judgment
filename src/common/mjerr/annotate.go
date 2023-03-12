package mjerr

import (
	"fmt"
)

type Annotator func(*mjError)

func WithOriginError(originError OriginError) Annotator {
	return func(err *mjError) {
		err.originError = originError
	}
}

func WithLogMessage(msg string) Annotator {
	return func(err *mjError) {
		err.logMessage = msg
	}
}

func WithLogMessagef(msg string, args ...interface{}) Annotator {
	return WithLogMessage(fmt.Sprintf(msg, args...))
}

func WithLogDetail(detail map[string]interface{}) Annotator {
	return func(err *mjError) {
		err.logDetail = detail
	}
}
