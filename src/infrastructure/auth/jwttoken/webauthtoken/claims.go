package webauthtoken

import (
	"crypto/subtle"
	"encoding/base64"
	"strconv"
	"time"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type webAuthTokenClaims struct {
	Issuer    string `json:"iss"`    // 発行者(subkarte固定文字列)
	Audience  string `json:"aud"`    // 利用者(base64 encode userID)
	UserID    int    `json:"userID"` // ログインユーザーID
	ExpiresAt int64  `json:"exp"`    // 有効期限(UNIX時間)
	IssuedAt  int64  `json:"iat"`    // 発行日時(UNIX時間)
}

func NewWebAuthTokenClaims(userID int, issuedAt time.Time) *webAuthTokenClaims {
	userIDStr := strconv.Itoa(userID)
	encodedOfficeID := base64.StdEncoding.EncodeToString([]byte(userIDStr))

	return &webAuthTokenClaims{
		Issuer:    webAuthTokenIssuer,
		Audience:  encodedOfficeID,
		UserID:    userID,
		ExpiresAt: issuedAt.Add(webAuthTokenExpireDuration).Unix(),
		IssuedAt:  issuedAt.Unix(),
	}
}

func (c *webAuthTokenClaims) OfficeID() int {
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

func (c *webAuthTokenClaims) Valid() error {
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

	if !c.verifyUserID() {
		return mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidToken),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"userID": c.UserID,
				},
			),
		)
	}

	if !c.verifyExpiresAt() {
		return mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidToken),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"expiresAt": time.Unix(c.ExpiresAt, 0).In(time.FixedZone("JST", 9*60*60)),
				},
			),
		)
	}

	return nil
}

func (c *webAuthTokenClaims) verifyIssuer() bool {
	if c.Issuer == "" {
		return false
	}

	if subtle.ConstantTimeCompare([]byte(c.Issuer), []byte(webAuthTokenIssuer)) == 0 {
		return false
	}

	return true
}

func (c *webAuthTokenClaims) verifyAudience() bool {
	return c.OfficeID() > 0
}

func (c *webAuthTokenClaims) verifyUserID() bool {
	return c.UserID > 0
}

func (c *webAuthTokenClaims) verifyExpiresAt() bool {
	return c.ExpiresAt >= time.Now().UTC().Unix()
}
