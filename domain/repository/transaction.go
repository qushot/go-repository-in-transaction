package repository

import "context"

type Transactioner interface {
	Begin(ctx context.Context) (context.Context, error)
	End(ctx context.Context, err error) (context.Context, error)
}
