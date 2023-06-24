package webclienttoken

import (
	"crypto/subtle"
	"encoding/base64"
	"strconv"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type webClientTokenClaims struct {
	Issuer   string `json:"iss"` // 発行者(固定文字列)
	Audience string `json:"aud"` // 利用者(クライアントがbase64 encodeしたuserID)
}

func (c *webClientTokenClaims) UserID() int {
	if c.Audience == "" {
		return 0
	}

	dec, err := base64.StdEncoding.DecodeString(c.Audience)
	if err != nil {
		return 0
	}

	userID, err := strconv.Atoi(string(dec))
	if err != nil {
		return 0
	}

	return userID
}

func (c *webClientTokenClaims) Valid() error {
	if !c.verifyIssuer() {
		return mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidToken),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"issuer": c.Issuer,
				},
			),
		)
	}

	if !c.verifyAudience() {
		return mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidToken),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"audience": c.Audience,
				},
			),
		)
	}

	return nil
}

func (c *webClientTokenClaims) verifyIssuer() bool {
	if c.Issuer == "" {
		return false
	}

	if subtle.ConstantTimeCompare([]byte(c.Issuer), []byte(webClientTokenIssuer)) == 0 {
		return false
	}

	return true
}

func (c *webClientTokenClaims) verifyAudience() bool {
	return c.UserID() > 0
}
