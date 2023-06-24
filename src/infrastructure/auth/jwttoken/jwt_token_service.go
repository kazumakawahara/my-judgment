package jwttoken

import (
	"time"

	"my-judgment/common/mjerr"
	"my-judgment/common/vo/uservo"
	"my-judgment/infrastructure/auth/jwttoken/webauthtoken"
	"my-judgment/infrastructure/auth/jwttoken/webclienttoken"
)

type jwtTokenService struct{}

func NewJwtTokenService() *jwtTokenService {
	return &jwtTokenService{}
}

func (s *jwtTokenService) ParseWebClientToken(webClientToken string) (userID int, err error) {
	webClaims, err := webclienttoken.ParseWebClientToken(webClientToken)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	return webClaims.UserID(), nil
}

func (s *jwtTokenService) GenerateWebAuthToken(userIDVO uservo.ID, issuedAt time.Time) (token string, err error) {
	claims := webauthtoken.NewWebAuthTokenClaims(userIDVO.Value(), issuedAt)

	token, err = webauthtoken.GenerateAuthToken(claims)
	if err != nil {
		return "", mjerr.Wrap(err)
	}

	return token, nil
}
