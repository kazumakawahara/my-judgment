package tokenhandler

import (
	"net/http"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/interfaces/handler"
	"my-judgment/interfaces/presenter"
	"my-judgment/usecase/tokenusecase"
	"my-judgment/usecase/tokenusecase/tokeninput"
)

type generateWebTokenHandler struct {
	generateTokenUsecase tokenusecase.GenerateWebTokenUsecase
}

func NewGenerateWebTokenHandler(generateTokenUsecase tokenusecase.GenerateWebTokenUsecase) *generateWebTokenHandler {
	return &generateWebTokenHandler{
		generateTokenUsecase: generateTokenUsecase,
	}
}

func (h *generateWebTokenHandler) GenerateWebToken(c handler.Context) error {
	clientToken := c.Request().Header.Get("Authorization")
	if clientToken == "" {
		return presenter.ErrorJSON(c, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.TokenRequired)))
	}

	in := &tokeninput.GenerateWebTokenInput{
		ClientToken: clientToken,
	}

	out, err := h.generateTokenUsecase.GenerateWebToken(c.Request().Context(), in)
	if err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err))
	}

	return presenter.JSON(c, http.StatusOK, out)
}
