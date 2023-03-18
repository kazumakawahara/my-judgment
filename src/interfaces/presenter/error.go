package presenter

import (
	"context"
	"log"

	"my-judgment/common/apperr"
	"my-judgment/common/config"
	"my-judgment/common/mjerr"
)

type httpError struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code   string        `json:"code"`
	Detail apperr.Detail `json:"detail"`
}

func ErrorJSON(c contextForPresenter, err error) error {
	r := c.Request()
	ctx := context.WithValue(r.Context(), config.LogKey, err)
	*r = *r.WithContext(ctx)

	// デフォルトエラーレスポンス
	statusCode := apperr.InternalServerError.StatusCode()
	httpErr := &httpError{
		Error: errorBody{
			Code:   apperr.InternalServerError.Error(),
			Detail: nil,
		},
	}

	mjErr := mjerr.AsApoError(err)
	if mjErr == nil {
		log.Print("failed in type assertion for mjError")
		return c.JSON(statusCode, httpErr)
	}

	appErr := mjErr.OriginError()
	if appErr == nil {
		log.Print("failed in type assertion for OriginError")
		return c.JSON(statusCode, httpErr)
	}

	apiErr := apperr.AsAPIError(appErr)
	if apiErr == nil {
		log.Print("failed in type assertion for apiError")
		return c.JSON(statusCode, httpErr)
	}

	statusCode = apiErr.StatusCode()
	httpErr = &httpError{
		Error: errorBody{
			Code:   apiErr.Error(),
			Detail: apiErr.Detail(),
		},
	}

	return c.JSON(statusCode, httpErr)
}
