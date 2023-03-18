package tx

import (
	"context"

	"my-judgment/common/apperr"
	mjerr2 "my-judgment/common/mjerr"
)

func HandleErrorWithRollbackTx(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if _, txErr := RollbackTx(ctx); txErr != nil {
		return mjerr2.Wrap(
			err,
			mjerr2.WithOriginError(apperr.InternalServerError),
			mjerr2.WithLogMessagef("(\n%s\n)", txErr.Error()),
		)
	}

	return mjerr2.Wrap(err)
}
