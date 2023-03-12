package uservo

import (
	"unicode/utf8"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Name string

func NewName(name string) (Name, error) {
	if l := utf8.RuneCountInString(name); l < 1 || l > 20 {
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"name": name,
				},
			),
		)
	}

	return Name(name), nil
}

func (t Name) Value() string {
	return string(t)
}
