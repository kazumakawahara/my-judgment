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

type updateUserHandler struct {
	updateUserUsecase userusecase.UpdateUserUsecase
}

func NewUpdateUserHandler(updateUserUsecase userusecase.UpdateUserUsecase) *updateUserHandler {
	return &updateUserHandler{
		updateUserUsecase: updateUserUsecase,
	}
}

func (h *updateUserHandler) UpdateUser(c handler.Context) error {
	in := &userinput.UpdateUserInput{}

	if err := c.Bind(in); err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err, mjerr.WithOriginError(apperr.InvalidParameter)))
	}

	out, err := h.updateUserUsecase.UpdateUser(c.Request().Context(), in)
	if err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err))
	}

	return presenter.JSON(c, http.StatusOK, out)
}
