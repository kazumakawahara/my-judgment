package useroutput

import (
	"my-judgment/common/vo/uservo"
)

type UpdateUserOutput struct {
	User UpdateUser `json:"mjUser"`
}

type UpdateUser struct {
	ID int `json:"id"`
}

func NewUpdateUserOutput(idVO uservo.ID) *UpdateUserOutput {
	return &UpdateUserOutput{
		User: UpdateUser{
			ID: idVO.Value(),
		},
	}
}
