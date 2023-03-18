package userdm

import (
	"time"

	"my-judgment/common/mjerr"
	"my-judgment/common/vo/sharedvo"
	"my-judgment/common/vo/uservo"
)

func Reconstruct(
	id int,
	name string,
	birthday time.Time,
	gender string,
	address string,
	email string,
	password string,
	plan int,
	createdAt time.Time,
	createdBy int,
	updatedAt time.Time,
	updatedBy int,
	deletedAt *time.Time,
	deletedBy *int,
) (*User, error) {
	idVO, err := uservo.NewID(id)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	nameVO, err := uservo.NewName(name)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	birthdayVO, err := sharedvo.NewAuditTime(birthday)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	genderVO, err := uservo.NewGender(gender)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	addressVO, err := uservo.NewAddress(address)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	emailVO, err := uservo.NewEmail(email)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	passwordVO, err := uservo.NewPassword(password)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	planVO, err := uservo.NewPlan(plan)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	createdAtVO, err := sharedvo.NewAuditTime(createdAt)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	createdByVO, err := uservo.NewID(createdBy)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	updatedAtVO, err := sharedvo.NewAuditTime(updatedAt)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	updatedByVO, err := uservo.NewID(updatedBy)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	deletedAtVO, err := sharedvo.NewNullableAuditTime(deletedAt)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	deletedByVO, err := uservo.NewNullableID(deletedBy)
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
		planVO,
		createdAtVO,
		createdByVO,
		updatedAtVO,
		updatedByVO,
		deletedAtVO,
		deletedByVO,
	)
}
