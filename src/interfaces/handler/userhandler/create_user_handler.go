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

type createUserHandler struct {
	createUserUsecase userusecase.CreateUserUsecase
}

func NewCreateUserHandler(createUserUsecase userusecase.CreateUserUsecase) *createUserHandler {
	return &createUserHandler{
		createUserUsecase: createUserUsecase,
	}
}

func (h *createUserHandler) CreateUser(c handler.Context) error {
	in := &userinput.CreateUserInput{}

	if err := c.Bind(in); err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err, mjerr.WithOriginError(apperr.InvalidParameter)))
	}

	out, err := h.createUserUsecase.CreateUser(c.Request().Context(), in)
	if err != nil {
		return presenter.ErrorJSON(c, mjerr.Wrap(err))
	}

	return presenter.JSON(c, http.StatusOK, out)
}
