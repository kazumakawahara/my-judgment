package uservo

import (
	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type ID int

func NewID(id int) (ID, error) {
	if id < 0 {
		return 0, mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"id": id,
				},
			),
		)
	}

	return ID(id), nil
}

func NewNotPersistedID(id int) (ID, error) {
	if id < 0 {
		return 0, mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"id": id,
				},
			),
		)
	}

	return ID(id), nil
}

func NewNullableID(id *int) (*ID, error) {
	if id == nil {
		return nil, nil
	}

	idVO, err := NewID(*id)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return &idVO, nil
}

func (i ID) Value() int {
	return int(i)
}

func (i *ID) NullableValue() *int {
	if i == nil {
		return nil
	}

	idVal := i.Value()

	return &idVal
}

func (i ID) Equals(id ID) bool {
	return i.Value() == id.Value()
}

const UserIDForNoUser ID = 0 // ユーザーなし時のID
