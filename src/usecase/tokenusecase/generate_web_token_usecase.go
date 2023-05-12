//go:generate mockgen -source=$GOFILE -destination=../../mock/mockusecase/mock$GOPACKAGE/mock_$GOFILE -package=mock$GOPACKAGE
package tokenusecase

import (
	"context"
	"time"

	"my-judgment/common/mjerr"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/usecase/tokenusecase/tokeninput"
	"my-judgment/usecase/tokenusecase/tokenoutput"
	"my-judgment/usecase/tokenusecase/tokenservice"
)

type GenerateWebTokenUsecase interface {
	GenerateWebToken(ctx context.Context, in *tokeninput.GenerateWebTokenInput) (*tokenoutput.GenerateWebTokenOutput, error)
}

type generateWebTokenUsecase struct {
	tokenService   tokenservice.TokenService
	userRepository userdm.Repository
}

func NewGenerateWebTokenUsecase(
	tokenService tokenservice.TokenService,
	userRepository userdm.Repository,
) *generateWebTokenUsecase {
	return &generateWebTokenUsecase{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (u *generateWebTokenUsecase) GenerateWebToken(ctx context.Context, in *tokeninput.GenerateWebTokenInput) (*tokenoutput.GenerateWebTokenOutput, error) {
	//userID, err := u.tokenService.ParseWebClientToken(in.ClientToken)
	//if err != nil {
	//	return nil, mjerr.Wrap(err)
	//}

	// TODO
	userID := 1000

	userIDVO, err := uservo.NewID(userID)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	if _, err := u.userRepository.FetchUserByID(ctx, userIDVO, false); err != nil {
		return nil, mjerr.Wrap(err)
	}

	token, err := u.tokenService.GenerateWebAuthToken(userIDVO, time.Now().UTC())
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return tokenoutput.NewGenerateWebTokenOutput(token), nil
}
