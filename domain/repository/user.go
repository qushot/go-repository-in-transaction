package repository

import (
	"context"

	"github.com/qushot/go-repository-in-transaction/domain"
)

type User interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
}
