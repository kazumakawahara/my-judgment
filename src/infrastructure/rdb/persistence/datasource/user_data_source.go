package datasource

import (
	"time"

	"my-judgment/common/mjerr"
	"my-judgment/domain/userdm"
)

type User struct {
	ID        int        `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	Birthday  time.Time  `gorm:"column:birthday"`
	Gender    string     `gorm:"column:gender"`
	Address   string     `gorm:"column:address"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	Plan      int        `gorm:"column:plan"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	CreatedBy int        `gorm:"column:created_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	UpdatedBy int        `gorm:"column:updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy *int       `gorm:"column:deleted_by"`
}

func (p *User) TableName() string {
	return "users"
}

func (p *User) ReconstructUserEntity() (*userdm.User, error) {
	userEntity, err := userdm.Reconstruct(
		p.ID,
		p.Name,
		p.Birthday,
		p.Gender,
		p.Address,
		p.Email,
		p.Password,
		p.Plan,
		p.CreatedAt,
		p.CreatedBy,
		p.UpdatedAt,
		p.UpdatedBy,
		p.DeletedAt,
		p.DeletedBy,
	)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}
	return userEntity, nil
}

func NewUser(userEntity *userdm.User) *User {
	return &User{
		ID:        userEntity.ID().Value(),
		Name:      userEntity.Name().Value(),
		Birthday:  userEntity.Birthday().Value(),
		Gender:    userEntity.Gender().Value(),
		Address:   userEntity.Address().Value(),
		Email:     userEntity.Email().Value(),
		Password:  userEntity.Password().Value(),
		Plan:      userEntity.Plan().Value(),
		CreatedAt: userEntity.CreatedAt().Value(),
		CreatedBy: userEntity.CreatedBy().Value(),
		UpdatedAt: userEntity.UpdatedAt().Value(),
		UpdatedBy: userEntity.UpdatedBy().Value(),
		DeletedAt: userEntity.DeletedAt().NullableValue(),
		DeletedBy: userEntity.DeletedBy().NullableValue(),
	}
}
