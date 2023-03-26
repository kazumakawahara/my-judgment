package useroutput

import (
	"time"

	"my-judgment/domain/userdm"
)

type FetchUserOutput struct {
	User FetchUser `json:"mjUser"`
}

type FetchUser struct {
	Name      string    `json:"name"`
	Birthday  time.Time `json:"birthday"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Plan      int       `json:"plan"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewFetchUserOutput(userEntity *userdm.User) *FetchUserOutput {
	return &FetchUserOutput{
		User: FetchUser{
			Name:      userEntity.Name().Value(),
			Birthday:  userEntity.Birthday().Value(),
			Gender:    userEntity.Gender().Value(),
			Address:   userEntity.Address().Value(),
			Email:     userEntity.Email().Value(),
			Plan:      userEntity.Plan().Value(),
			CreatedAt: userEntity.CreatedAt().Value(),
		},
	}
}
