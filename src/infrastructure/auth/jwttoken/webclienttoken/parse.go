package webclienttoken

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

func ParseWebClientToken(token string) (*webClientTokenClaims, error) {
	claims := &webClientTokenClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv(webClientTokenJwtPrivateKey)
		return []byte(secretKey), nil
	}

	signingMethodOption := jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()})
	parser := jwt.NewParser(signingMethodOption)
	if _, err := parser.ParseWithClaims(token, claims, keyFunc); err != nil {
		if errors.Is(err, apperr.InvalidToken) {
			return nil, mjerr.Wrap(err)
		}

		return nil, mjerr.Wrap(err, mjerr.WithOriginError(apperr.InvalidRequestToken))
	}

	return claims, nil
}
