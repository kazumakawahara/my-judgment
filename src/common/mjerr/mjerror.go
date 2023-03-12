package mjerr

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/xerrors"
)

type mjError struct {
	// Error for log output
	next       error
	logMessage string
	logDetail  map[string]interface{}
	errorTime  time.Time
	stackTrace stackTrace

	// Error for app
	originError OriginError
}

const (
	levelFormat      = "[level]: %s\n"
	codeFormat       = "[code]: %s\n"
	messageFormat    = "[message]: %s\n"
	detailFormat     = "[detail]: %s\n"
	errorTimeFormat  = "[errorTime]: %s\n"
	stackTraceFormat = "[stackTrace]: %s"
)

func (e *mjError) Error() string {
	var message string

	if e.originError == nil {
		message += fmt.Sprintf(levelFormat, "-")
		message += fmt.Sprintf(codeFormat, "-")
	} else {
		message += fmt.Sprintf(levelFormat, e.originError.Level())
		message += fmt.Sprintf(codeFormat, e.originError.Error())
	}

	if e.logMessage == "" {
		message += fmt.Sprintf(messageFormat, "-")
	} else {
		message += fmt.Sprintf(messageFormat, e.logMessage)
	}

	if len(e.logDetail) == 0 {
		message += fmt.Sprintf(detailFormat, "-")
	} else {
		b, err := json.Marshal(e.logDetail)
		if err != nil {
			log.Printf("mjerr package json.Marshal error: %+v", e.logDetail)
		}

		message += fmt.Sprintf(detailFormat, string(b))
	}

	message += fmt.Sprintf(errorTimeFormat, e.errorTime.Format(time.RFC3339Nano))
	message += fmt.Sprintf(stackTraceFormat, e.stackTrace.format())

	return message
}

func (e *mjError) OriginError() OriginError {
	if e == nil {
		return nil
	}

	if e.originError != nil {
		return e.originError
	}

	next := AsApoError(e.next)

	return next.OriginError()
}

func (e *mjError) Is(err error) bool {
	if target, ok := err.(OriginError); ok {
		if e.originError != nil &&
			e.originError.Error() == target.Error() {
			return true
		}

		next := AsApoError(e.next)
		if next == nil {
			return false
		}

		return next.Is(err)
	}

	return false
}

func (e *mjError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *mjError) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())

	return e.next
}

func New(msg string) error {
	return &mjError{
		logMessage: msg,
		errorTime:  time.Now().In(time.FixedZone("JST", 9*60*60)),
		stackTrace: newStackTrace(),
	}
}

func Errorf(format string, args ...interface{}) error {
	return &mjError{
		logMessage: fmt.Sprintf(format, args...),
		errorTime:  time.Now().In(time.FixedZone("JST", 9*60*60)),
		stackTrace: newStackTrace(),
	}
}

func Wrap(err error, annotators ...Annotator) error {
	apoErr := &mjError{
		next:       err,
		errorTime:  time.Now().In(time.FixedZone("JST", 9*60*60)),
		stackTrace: newStackTrace(),
	}

	for _, f := range annotators {
		f(apoErr)
	}

	return apoErr
}

func AsApoError(err error) *mjError {
	var e *mjError
	if errors.As(err, &e) {
		return e
	}

	return nil
}
