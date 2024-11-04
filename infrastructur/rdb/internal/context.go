package internal

import (
	"context"
	"database/sql"
)

type ctxTxKey struct{}

func WithTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, ctxTxKey{}, tx)
}

func WithoutTx(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxTxKey{}, nil)
}

func TxFromContext(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(ctxTxKey{}).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}
