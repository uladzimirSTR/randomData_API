package dbase

import (
	"context"
	"fmt"
	"log"
	"strings"

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

	sqlText := strings.TrimSpace(string(query))

	log.Printf("executing: %s", sqlText)
	if sqlText == "" {
		return fmt.Errorf("trim query: %w", err)
	}

	if _, err := pool.Exec(ctx, sqlText); err != nil {
		return fmt.Errorf("execute %s: %w", sqlText, err)
	}

	log.Println("all sql files executed successfully")
	return nil
}
