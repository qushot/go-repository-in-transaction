package rdb

import (
	"context"
	"database/sql"

	"github.com/qushot/go-repository-in-transaction/domain"
	"github.com/qushot/go-repository-in-transaction/domain/repository"
	"github.com/qushot/go-repository-in-transaction/infrastructur/rdb/internal"
)

type UserRepository struct {
	db *sql.DB
}

// Create implements repository.User.
func (u *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (id, name, address) VALUES ($1, $2, $3)"

	exec := u.db.Exec
	tx := internal.TxFromContext(ctx)
	if tx != nil {
		exec = tx.Exec
	}

	if _, err := exec(query, user.ID, user.Name, user.Address); err != nil {
		return err
	}

	return nil
}

// Update implements repository.User.
func (u *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET name = $1, address = $2 WHERE id = $3"

	exec := u.db.Exec
	tx := internal.TxFromContext(ctx)
	if tx != nil {
		exec = tx.Exec
	}

	if _, err := exec(query, user.Name, user.Address, user.ID); err != nil {
		return err
	}

	return nil
}

// Delete implements repository.User.
func (u *UserRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"

	exec := u.db.Exec
	tx := internal.TxFromContext(ctx)
	if tx != nil {
		exec = tx.Exec
	}

	if _, err := exec(query, id); err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) repository.User {
	return &UserRepository{db: db}
}
