package sharedvo

import (
	"time"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type AuditTime time.Time

func NewAuditTime(date time.Time) (AuditTime, error) {
	if date.IsZero() {
		return AuditTime(time.Time{}), mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InvalidParameter))
	}

	return AuditTime(date), nil
}

func (d AuditTime) Value() time.Time {
	return time.Time(d)
}

func NewNullableAuditTime(dateTime *time.Time) (*AuditTime, error) {
	if dateTime == nil {
		return nil, nil
	}

	dateTimeVO, err := NewAuditTime(*dateTime)
	if err != nil {
		return nil, mjerr.Wrap(err)
	}

	return &dateTimeVO, nil
}

func (d *AuditTime) NullableValue() *time.Time {
	if d == nil {
		return nil
	}

	dateTime := time.Time(*d)

	return &dateTime
}

func (d *AuditTime) UnixNanoTimestamp() int {
	if d == nil {
		return 0
	}

	return int(d.NullableValue().UnixNano())
}
