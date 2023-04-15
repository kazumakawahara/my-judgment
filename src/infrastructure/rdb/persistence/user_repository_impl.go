package persistence

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/infrastructure/rdb"
	"my-judgment/infrastructure/rdb/persistence/datasource"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, userEntity *userdm.User) (uservo.ID, error) {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	userDS := datasource.NewUser(userEntity)

	if res := conn.Create(userDS); res.Error != nil {
		return 0, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.InternalServerError))
	}

	idVO, err := uservo.NewID(userDS.ID)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	return idVO, nil
}

func (r *userRepository) ExistsUserByPassword(ctx context.Context, passwordVO uservo.Password) (bool, error) {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return false, mjerr.Wrap(err)
	}

	var count int64

	if res := conn.
		Model(&datasource.User{}).
		Where("password = ?", passwordVO.Value()).
		Where("deleted_at IS NULL").
		Count(&count); res.Error != nil {
		return false, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))
	}

	return count > 0, nil
}

func (r *userRepository) FetchUserByID(ctx context.Context, userIDVO uservo.ID, withLock bool) (*userdm.User, error) {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	userDS := datasource.User{}

	query := conn.
		Scopes(scopeForUser()).
		Where("id = ?", userIDVO.Value())

	if withLock {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	if res := query.Take(&userDS); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.MjUserNotFound))
		}

		return nil, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.InternalServerError))
	}

	userEntity, err := userDS.ReconstructUserEntity()
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return userEntity, nil
}

func (r *userRepository) FetchUserIDByName(ctx context.Context, nameVO uservo.Name) (uservo.ID, error) {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	var id int

	if res := conn.
		Select("id").
		Scopes(scopeForUser()).
		Where("name = ?", nameVO.Value()).
		Take(&id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return 0, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.MjUserNotFound))
		}

		return 0, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.InternalServerError))
	}

	idVO, err := uservo.NewID(id)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	return idVO, nil
}

func (r *userRepository) FetchUserIDByEmail(ctx context.Context, emailVO uservo.Email) (uservo.ID, error) {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	var id int

	if res := conn.
		Select("id").
		Scopes(scopeForUser()).
		Where("email = ?", emailVO.Value()).
		Take(&id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return 0, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.MjUserNotFound))
		}

		return 0, mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.InternalServerError))
	}

	idVO, err := uservo.NewID(id)
	if err != nil {
		return 0, mjerr.Wrap(err)
	}

	return idVO, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userEntity *userdm.User) error {
	conn, err := rdb.DBConnFromCtx(ctx)
	if err != nil {
		return mjerr.Wrap(err)
	}

	userDS := &datasource.User{
		Name:      userEntity.Name().Value(),
		Gender:    userEntity.Gender().Value(),
		Address:   userEntity.Address().Value(),
		Email:     userEntity.Email().Value(),
		Password:  userEntity.Password().Value(),
		UpdatedAt: userEntity.UpdatedAt().Value(),
		UpdatedBy: userEntity.UpdatedBy().Value(),
	}

	if res := conn.
		Scopes(scopeForUser()).
		Select("name", "gender", "address", "email", "password", "updated_at", "updated_by").
		Where("id = ?", userEntity.ID().Value()).
		Updates(userDS); res.Error != nil {
		return mjerr.Wrap(res.Error, mjerr.WithOriginError(apperr.InternalServerError))
	}

	return nil
}

func scopeForUser() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Model(&datasource.User{}).
			Where("deleted_at IS NULL").
			Where("deleted_uts = 0")
	}
}
