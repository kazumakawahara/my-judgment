package uservo

import (
	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Plan int

func NewPlan(plan int) (Plan, error) {
	if plan < 1 {
		return 0, mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"plan": plan,
				},
			),
		)
	}

	return Plan(plan), nil
}

func NewNotPersistedPlan(plan int) (Plan, error) {
	if plan < 0 {
		return 0, mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"plan": plan,
				},
			),
		)
	}

	return Plan(plan), nil
}

func (i Plan) Value() int {
	return int(i)
}

func (i Plan) Equals(plan Plan) bool {
	return i.Value() == plan.Value()
}

const (
	FreePlan Plan = 0
)
