package apperr

var (
	// 400
	InvalidParameter = newBadRequest("InvalidParameter")

	// 401
	TokenRequired       = newUnauthorized("TokenRequired")
	InvalidToken        = newUnauthorized("InvalidToken")
	InvalidRequestToken = newUnauthorized("InvalidRequest")

	// 404
	MjUserNotFound = newNotFound("MjUserNotFound")

	// 409
	MjUserNameConflict  = newConflict("MjUserNameConflict")
	MjUserEmailConflict = newConflict("MjUserEmailConflict")

	// 410
	Gone = newGone("Gone")

	// 500
	InternalServerError = newInternalServerError("InternalServerError")
)
