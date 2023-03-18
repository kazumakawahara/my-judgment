package uservo

import (
	"regexp"
	"unicode/utf8"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Email string

var emailFormat = regexp.MustCompile(`^[A-Za-z0-9\.\+_-]+@+[A-Za-z0-9\.\+_-]+\.+[A-Za-z0-9\.\+_-]`)

func NewEmail(email string) (Email, error) {
	if l := utf8.RuneCountInString(email); l < 1 || l > 255 {
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"email": email,
				},
			),
		)
	}

	if !emailFormat.MatchString(email) {
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"email": email,
				},
			),
		)
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}
