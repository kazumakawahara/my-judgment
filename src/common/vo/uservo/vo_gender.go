package uservo

import (
	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Gender string

func NewGender(gender string) (Gender, error) {
	switch gender {
	case Gender00101.Value():
	case Gender00102.Value():
	case Gender00103.Value():
	default:
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"assetType": gender,
				},
			),
		)
	}

	return Gender(gender), nil
}

func (gender Gender) Value() string {
	return string(gender)
}

func (gender Gender) Equals(assetType Gender) bool {
	return gender.Value() == assetType.Value()
}

const (
	Gender00101 Gender = "00101" // 男性
	Gender00102 Gender = "00102" // 女性
	Gender00103 Gender = "00103" // 不明
)
