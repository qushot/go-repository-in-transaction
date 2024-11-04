package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/qushot/go-repository-in-transaction/domain"
	"github.com/qushot/go-repository-in-transaction/infrastructur/rdb"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user password=pass dbname=sampledb sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// トランザクションを使わない場合
	user1 := &domain.User{
		ID:      100,
		Name:    "Suzuki",
		Address: "Saitama",
	}

	userRepository1 := rdb.NewUserRepository(db)
	if err := userRepository1.Create(ctx, user1); err != nil {
		log.Printf("failed to create user1: %v", err)
		return
	}
	if err := userRepository1.Update(ctx, user1); err != nil {
		log.Printf("failed to update user1: %v", err)
		return
	}
	if err := userRepository1.Delete(ctx, 100); err != nil {
		log.Printf("failed to delete user1: %v", err)
		return
	}

	// トランザクションを使う場合
	user2 := &domain.User{
		ID:      200,
		Name:    "Takahashi",
		Address: "Osaka",
	}

	atomic := rdb.NewTransaction(db)
	ctx, err = atomic.Begin(ctx)
	if err != nil {
		log.Printf("failed to begin tx: %v", err)
		return
	}

	defer func() {
		if _, err := atomic.End(ctx, err); err != nil {
			log.Printf("failed to end tx: %v", err)
		}
	}()

	userRepository2 := rdb.NewUserRepository(db)
	if err := userRepository2.Create(ctx, user2); err != nil {
		log.Printf("failed to create user2: %v", err)
		return
	}
	if err := userRepository2.Update(ctx, user2); err != nil {
		log.Printf("failed to update user2: %v", err)
		return
	}
	if err := userRepository2.Delete(ctx, 200); err != nil {
		log.Printf("failed to delete user2: %v", err)
		return
	}
}
