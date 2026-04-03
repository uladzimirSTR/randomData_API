package dbase

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ExecuteSQL(ctx context.Context, query string) error {
	pool, err := pgxpool.New(ctx, DB_URL)

	if err != nil {
		return fmt.Errorf("create pg pool: %w", err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	log.Printf("executing: %s", query)

	if _, err := pool.Exec(ctx, query); err != nil {
		return fmt.Errorf("execute %s: %w", query, err)
	}

	log.Println("all sql files executed successfully")
	return nil
}
