package helper

import (
	"context"

	"gorm.io/gorm"
)

type txContextKey struct{}

var ctxTxKey = txContextKey{}

func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxTxKey, tx)
}

func GetTx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(ctxTxKey).(*gorm.DB)
	if ok {
		return tx
	}
	return db
}
