package webauthtoken

import (
	"os"

	"github.com/golang-jwt/jwt/v4"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

func GenerateAuthToken(claims *webAuthTokenClaims) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(os.Getenv(webAuthTokenJwtPrivateKey)))
	if err != nil {
		return "", mjerr.Wrap(err, mjerr.WithOriginError(apperr.InternalServerError))
	}

	return token, nil
}
