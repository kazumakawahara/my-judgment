//go:generate mockgen -source=$GOFILE -destination=../../mock/mockrepository/mockuserrepository/mock_user_$GOFILE -package=mockuserrepository
package userdm

import (
	"context"

	"my-judgment/common/vo/uservo"
)

type Repository interface {
	CreateUser(ctx context.Context, userEntity *User) (uservo.ID, error)
	ExistsUserByPassword(ctx context.Context, passwordVO uservo.Password) (bool, error)
	FetchUserByID(ctx context.Context, userIDVO uservo.ID, withLock bool) (*User, error)
	FetchUserIDByName(ctx context.Context, nameVO uservo.Name) (uservo.ID, error)
	FetchUserIDByEmail(ctx context.Context, emailVO uservo.Email) (uservo.ID, error)
	UpdateUser(ctx context.Context, userEntity *User) error
}
