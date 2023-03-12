package userdm

import (
	"my-judgment/common/vo/sharedvo"
	"my-judgment/common/vo/uservo"
)

type User struct {
	id        uservo.ID
	name      uservo.Name
	birthday  sharedvo.AuditTime
	gender    uservo.Gender
	address   uservo.Address
	email     uservo.Email
	password  uservo.Password
	plan      uservo.Plan
	createdAt sharedvo.AuditTime
	createdBy uservo.ID
	updatedAt sharedvo.AuditTime
	updatedBy uservo.ID
	deletedAt *sharedvo.AuditTime
	deletedBy *uservo.ID
}

func newUser(
	idVO uservo.ID,
	nameVO uservo.Name,
	birthdayVO sharedvo.AuditTime,
	genderVO uservo.Gender,
	addressVO uservo.Address,
	emailVO uservo.Email,
	passwordVO uservo.Password,
	planVO uservo.Plan,
	createdAtVO sharedvo.AuditTime,
	createdByVO uservo.ID,
	updatedAtVO sharedvo.AuditTime,
	updatedByVO uservo.ID,
	deletedAtVO *sharedvo.AuditTime,
	deletedByVO *uservo.ID,
) (*User, error) {
	return &User{
		id:        idVO,
		name:      nameVO,
		birthday:  birthdayVO,
		gender:    genderVO,
		address:   addressVO,
		email:     emailVO,
		password:  passwordVO,
		plan:      planVO,
		createdAt: createdAtVO,
		createdBy: createdByVO,
		updatedAt: updatedAtVO,
		updatedBy: updatedByVO,
		deletedAt: deletedAtVO,
		deletedBy: deletedByVO,
	}, nil
}

func (u *User) ID() uservo.ID {
	return u.id
}

func (u *User) Name() uservo.Name {
	return u.name
}

func (u *User) Birthday() sharedvo.AuditTime {
	return u.birthday
}

func (u *User) Gender() uservo.Gender {
	return u.gender
}

func (u *User) Address() uservo.Address {
	return u.address
}

func (u *User) Email() uservo.Email {
	return u.email
}

func (u *User) Password() uservo.Password {
	return u.password
}

func (u *User) Plan() uservo.Plan {
	return u.plan
}

func (u *User) CreatedAt() sharedvo.AuditTime {
	return u.createdAt
}

func (u *User) CreatedBy() uservo.ID {
	return u.createdBy
}

func (u *User) UpdatedAt() sharedvo.AuditTime {
	return u.updatedAt
}

func (u *User) UpdatedBy() uservo.ID {
	return u.updatedBy
}

func (u *User) DeletedAt() *sharedvo.AuditTime {
	return u.deletedAt
}

func (u *User) DeletedBy() *uservo.ID {
	return u.deletedBy
}
