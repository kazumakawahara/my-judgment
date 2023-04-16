//go:generate mockgen -source=$GOFILE -destination=../../mock/mockusecase/mock$GOPACKAGE/mock_$GOFILE -package=mock$GOPACKAGE
package userusecase

import (
	"context"

	"my-judgment/common/mjerr"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

type FetchUserUsecase interface {
	FetchUser(ctx context.Context, in *userinput.FetchUserInput) (*useroutput.FetchUserOutput, error)
}

type fetchUserUsecase struct {
	userRepository userdm.Repository
}

func NewFetchUserUsecase(userRepository userdm.Repository) *fetchUserUsecase {
	return &fetchUserUsecase{
		userRepository: userRepository,
	}
}

func (u *fetchUserUsecase) FetchUser(ctx context.Context, in *userinput.FetchUserInput) (*useroutput.FetchUserOutput, error) {
	userIDVO, err := uservo.NewID(in.UserID)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	userEntity, err := u.userRepository.FetchUserByID(ctx, userIDVO, false)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return useroutput.NewFetchUserOutput(userEntity), nil
}
