package uservo

import (
	"math/rand"
	"time"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Password string

func NewPassword(password string) (Password, error) {
	if l := len(password); l != 16 {
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"password": password,
				},
			),
		)
	}

	return Password(password), nil
}

func (n Password) Value() string {
	return string(n)
}

const (
	rs6Letters       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

func GeneratePassword(n int) (Password, error) {
	randSrc := rand.NewSource(time.Now().UTC().UnixNano())

	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}

	passwordVO, err := NewPassword(string(b))
	if err != nil {
		return "", mjerr.Wrap(err)
	}

	return passwordVO, nil
}
