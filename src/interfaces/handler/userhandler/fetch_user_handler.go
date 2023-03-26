package userhandler

import (
	"net/http"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/interfaces/handler"
	"my-judgment/interfaces/presenter"
	"my-judgment/usecase/userusecase"
	"my-judgment/usecase/userusecase/userinput"
)

type fetchUserHandler struct {
	fetchUserUsecase userusecase.FetchUserUsecase
}

func NewFetchUserHandler(fetchUserUsecase userusecase.FetchUserUsecase) *fetchUserHandler {
	return &fetchUserHandler{
		fetchUserUsecase: fetchUserUsecase,
	}
}

func (h *fetchUserHandler) FetchUser(c handler.Context) error {
	in := &userinput.FetchUserInput{}

	if err := c.Bind(in); err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err, mjerr.WithOriginError(apperr.InvalidParameter)))
	}

	out, err := h.fetchUserUsecase.FetchUser(c.Request().Context(), in)
	if err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err))
	}

	return presenter.JSON(c, http.StatusOK, out)
}
