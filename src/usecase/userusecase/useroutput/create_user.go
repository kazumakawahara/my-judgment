package useroutput

import (
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
)

type CreateUserOutput struct {
	User CreateUser `json:"mjUser"`
}

type CreateUser struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func NewCreateUserOutput(idVO uservo.ID, userEntity *userdm.User) *CreateUserOutput {
	return &CreateUserOutput{
		User: CreateUser{
			ID:       idVO.Value(),
			Password: userEntity.Password().Value(),
		},
	}
}
