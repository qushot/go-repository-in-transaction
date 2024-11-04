package rdb

import (
	"context"
	"database/sql"

	"github.com/qushot/go-repository-in-transaction/domain/repository"
	"github.com/qushot/go-repository-in-transaction/infrastructur/rdb/internal"
)

type Transaction struct {
	db *sql.DB
}

func (t *Transaction) Begin(ctx context.Context) (context.Context, error) {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return ctx, err
	}

	return internal.WithTx(ctx, tx), nil
}

func (t *Transaction) End(ctx context.Context, err error) (context.Context, error) {
	tx := internal.TxFromContext(ctx)
	if tx == nil {
		return ctx, nil
	}

	withoutCtx := internal.WithoutTx(ctx)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return withoutCtx, err
		}
		return withoutCtx, nil
	}

	if err := tx.Commit(); err != nil {
		return withoutCtx, err
	}

	return withoutCtx, nil
}

// こいつが tx を含んだ ctx を返すようにして、End 的なメソッドもこいつが返したほうがいいかも。
func NewTransaction(db *sql.DB) repository.Transactioner {
	return &Transaction{db: db}
}
