package rdb

import (
	"context"

	"gorm.io/gorm"

	"my-judgment/common/apperr"
	"my-judgment/common/config"
	"my-judgment/common/mjerr"
)

func DBConnFromCtx(ctx context.Context) (*gorm.DB, error) {
	txVal := ctx.Value(config.TXKey)
	if txVal != nil {
		tx, ok := txVal.(*gorm.DB)
		if !ok {
			return nil, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))
		}

		return tx, nil
	}

	connVal := ctx.Value(config.DBKey)
	if connVal == nil {
		return nil, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))
	}

	conn, ok := connVal.(*gorm.DB)
	if !ok {
		return nil, mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))
	}

	return conn, nil
}
