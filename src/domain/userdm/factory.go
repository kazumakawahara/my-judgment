package userdm

import (
	"time"

	"my-judgment/common/mjerr"
	"my-judgment/common/vo/sharedvo"
	"my-judgment/common/vo/uservo"
)

func GenUserForCreate(
	nameVO uservo.Name,
	birthdayVO sharedvo.AuditTime,
	genderVO uservo.Gender,
	addressVO uservo.Address,
	emailVO uservo.Email,
	passwordVO uservo.Password,
) (*User, error) {
	// DBレコード作成時にIDは採番されるため0値
	idVO, err := uservo.NewNotPersistedID(0)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	now := time.Now().UTC()

	createdAtVO, err := sharedvo.NewAuditTime(now)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	updatedAtVO, err := sharedvo.NewAuditTime(now)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return newUser(
		idVO,
		nameVO,
		birthdayVO,
		genderVO,
		addressVO,
		emailVO,
		passwordVO,
		uservo.FreePlan,
		createdAtVO,
		uservo.UserIDForNoUser,
		updatedAtVO,
		uservo.UserIDForNoUser,
		nil,
		nil,
	)
}
