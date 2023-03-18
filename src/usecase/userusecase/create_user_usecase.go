package userusecase

import (
	"context"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/tx"
	"my-judgment/common/vo/sharedvo"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

type CreateUserUsecase interface {
	CreateUser(ctx context.Context, in *userinput.CreateUserInput) (*useroutput.CreateUserOutput, error)
}

type createUserUsecase struct {
	userRepository userdm.Repository
}

func NewCreateUserUsecase(userRepository userdm.Repository) *createUserUsecase {
	return &createUserUsecase{
		userRepository: userRepository,
	}
}

func (u *createUserUsecase) CreateUser(ctx context.Context, in *userinput.CreateUserInput) (*useroutput.CreateUserOutput, error) {
	nameVO, err := uservo.NewName(in.User.Name)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	birthdayVO, err := sharedvo.NewAuditTime(in.User.Birthday)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	genderVO, err := uservo.NewGender(in.User.Gender)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	addressVO, err := uservo.NewAddress(in.User.Address)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	emailVO, err := uservo.NewEmail(in.User.Email)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	passwordVO, err := u.CheckPassword(ctx)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	userEntity, err := userdm.GenUserForCreate(
		nameVO,
		birthdayVO,
		genderVO,
		addressVO,
		emailVO,
		passwordVO,
	)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	ctx, rollback, err := tx.BeginTx(ctx)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}
	defer rollback()

	userDomainService := userdm.NewUserDomainService(u.userRepository)

	// ユーザー名重複確認
	if exist, err := userDomainService.ExistsUserByNameForCreate(ctx, nameVO); err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	} else if exist {
		err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict))
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	// Eメール重複確認
	if exist, err := userDomainService.ExistsUserByEmailForCreate(ctx, emailVO); err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	} else if exist {
		err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserEmailConflict))
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	idVO, err := u.userRepository.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	ctx, err = tx.CommitTx(ctx)
	if err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	out := useroutput.NewCreateUserOutput(idVO, userEntity)

	return out, nil
}

func (u *createUserUsecase) CheckPassword(ctx context.Context) (uservo.Password, error) {
	var passwordVO uservo.Password

	for {
		// password新規生成
		generatePasswordVO, err := uservo.GeneratePassword(16)
		if err != nil {
			return "", mjerr.Wrap(err)
		}

		// password重複確認
		if exist, err := u.userRepository.ExistsUserByPassword(ctx, generatePasswordVO); err != nil {
			return "", mjerr.Wrap(err)
		} else if !exist {
			passwordVO = generatePasswordVO
			break
		} else {
			continue
		}
	}

	return passwordVO, nil
}
