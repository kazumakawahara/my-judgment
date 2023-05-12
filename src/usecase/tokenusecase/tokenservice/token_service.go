//go:generate mockgen -source=$GOFILE -destination=../../../mock/mocktokenservice/mocktokenservice/mock_token_$GOFILE -package=mocktokenservice
package tokenservice

import (
	"time"

	"my-judgment/common/vo/uservo"
)

type TokenService interface {
	ParseWebClientToken(webClientToken string) (userID int, err error)
	GenerateWebAuthToken(userIDVO uservo.ID, issuedAt time.Time) (token string, err error)
}
