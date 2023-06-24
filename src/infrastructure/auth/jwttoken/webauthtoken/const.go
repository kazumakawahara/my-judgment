package webauthtoken

import "time"

const (
	webAuthTokenIssuer         string        = "myJudgment"
	webAuthTokenJwtPrivateKey  string        = "MY_JUDGMENT_TOKEN_JWT_KEY" //#nosec G101 - This is a false positive
	webAuthTokenExpireDuration time.Duration = 1 * time.Hour
)
