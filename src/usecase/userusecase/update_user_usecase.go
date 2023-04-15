package userusecase

import (
	"context"
	"time"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/tx"
	"my-judgment/common/vo/sharedvo"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

type UpdateUserUsecase interface {
	UpdateUser(ctx context.Context, in *userinput.UpdateUserInput) (*useroutput.UpdateUserOutput, error)
}

type updateUserUsecase struct {
	userRepository userdm.Repository
}

func NewUpdateUserUsecase(userRepository userdm.Repository) *updateUserUsecase {
	return &updateUserUsecase{
		userRepository: userRepository,
	}
}

func (u *updateUserUsecase) UpdateUser(ctx context.Context, in *userinput.UpdateUserInput) (*useroutput.UpdateUserOutput, error) {
	// ユーザー情報の更新要素が全てnilの場合はエラー
	if in.User.Name == nil &&
		in.User.Gender == nil &&
		in.User.Address == nil &&
		in.User.Email == nil &&
		in.User.Password == nil {
		return nil, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InvalidParameter))
	}

	idVO, err := uservo.NewID(in.UserID)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	now := time.Now().UTC()
	timeNowVO, err := sharedvo.NewAuditTime(now)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	// ユーザー存在確認
	userEntity, err := u.userRepository.FetchUserByID(ctx, idVO, true)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	ctx, rollback, err := tx.BeginTx(ctx)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}
	defer rollback()

	userDomainService := userdm.NewUserDomainService(u.userRepository)

	if in.User.Name != nil {
		nameVO, err := uservo.NewName(*in.User.Name)
		if err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		// ユーザー名重複確認
		if exist, err := userDomainService.ExistsUserByNameForUpdate(ctx, idVO, nameVO); err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		} else if exist {
			err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict))
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		userEntity.ChangeName(nameVO)
	}

	if in.User.Email != nil {
		emailVO, err := uservo.NewEmail(*in.User.Email)
		if err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		// Eメール重複確認
		if exist, err := userDomainService.ExistsUserByEmailForUpdate(ctx, idVO, emailVO); err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		} else if exist {
			err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserEmailConflict))
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		userEntity.ChangeEmail(emailVO)
	}

	if in.User.Gender != nil {
		genderVO, err := uservo.NewGender(*in.User.Gender)
		if err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		userEntity.ChangeGender(genderVO)
	}

	if in.User.Address != nil {
		addressVO, err := uservo.NewAddress(*in.User.Address)
		if err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		userEntity.ChangeAddress(addressVO)
	}

	if in.User.Password != nil {
		passwordVO, err := uservo.NewPassword(*in.User.Password)
		if err != nil {
			return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
		}

		userEntity.ChangePassword(passwordVO)
	}

	userEntity.ChangeUpdatedAt(timeNowVO)
	userEntity.ChangeUpdatedBy(idVO)

	if err = u.userRepository.UpdateUser(ctx, userEntity); err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	ctx, err = tx.CommitTx(ctx)
	if err != nil {
		return nil, mjerr.Wrap(tx.HandleErrorWithRollbackTx(ctx, err))
	}

	out := useroutput.NewUpdateUserOutput(idVO)

	return out, nil
}
