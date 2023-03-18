package tx

import (
	"context"

	"gorm.io/gorm"

	"my-judgment/common/apperr"
	"my-judgment/common/config"
	mjerr2 "my-judgment/common/mjerr"
)

func BeginTx(ctx context.Context) (context.Context, func(), error) {
	val := ctx.Value(config.TXKey)
	if val != nil {
		// トランザクションのネストは許容しない
		// 既にtxがあった場合、Rollbackしてエラーを返す
		tx, ok := val.(*gorm.DB)
		if !ok {
			return nil, nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
		}

		if err := tx.Rollback().Error; err != nil {
			return nil, nil, mjerr2.Wrap(err, mjerr2.WithOriginError(apperr.InternalServerError))
		}

		return nil, nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	val = ctx.Value(config.DBKey)
	if val == nil {
		return nil, nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	conn, ok := val.(*gorm.DB)
	if !ok {
		return nil, nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	tx := conn.Begin()
	if err := tx.Error; err != nil {
		return nil, nil, mjerr2.Wrap(tx.Error, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	ctx = context.WithValue(ctx, config.TXKey, tx)
	rollbackFn := func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}

	return ctx, rollbackFn, nil
}

func CommitTx(ctx context.Context) (context.Context, error) {
	val := ctx.Value(config.TXKey)
	if val == nil {
		return nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	tx, ok := val.(*gorm.DB)
	if !ok {
		return nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	if err := tx.Commit().Error; err != nil {
		return ctx, mjerr2.Wrap(err, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	ctx = context.WithValue(ctx, config.TXKey, nil)

	return ctx, nil
}

func RollbackTx(ctx context.Context) (context.Context, error) {
	val := ctx.Value(config.TXKey)
	if val == nil {
		return nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	tx, ok := val.(*gorm.DB)
	if !ok {
		return nil, mjerr2.Wrap(nil, mjerr2.WithOriginError(apperr.InternalServerError))
	}

	if err := tx.Rollback().Error; err != nil {
		return nil, mjerr2.Wrap(
			nil,
			mjerr2.WithOriginError(apperr.InternalServerError),
			mjerr2.WithLogMessage(err.Error()),
		)
	}

	ctx = context.WithValue(ctx, config.TXKey, nil)

	return ctx, nil
}
