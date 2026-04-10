package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/uladzimirSTR/randomData_API/dbase"
	rnd "github.com/uladzimirSTR/randomData_API/randomData"
)

const TIME_WORK = 5 * time.Minute
const EVERY_N_HOUR_GENERATE = 2 * time.Hour

func main() {
	// Create a context with a timeout to ensure the function doesn't run indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), TIME_WORK)
	defer cancel()
	// Create a connection pool to the database
	pool, err := pgxpool.New(ctx, db.DB_URL)
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("ping database: %v", err)
	}

	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	// Generate random data and insert it into the database
	go func() {
		// Generate random data every N hours
		for {
			go rnd.RandomDataUsers(pool, false)
			time.Sleep(EVERY_N_HOUR_GENERATE)
		}
	}()

	data, err := db.GetUsers(
		pool,
		"random_data",
		"users",
		map[string]string{
			"dateCol": "updated_at",
			"start":   "2026-04-04",
			"end":     "2026-04-04",
		},
	)

	fmt.Printf("data: %+v\n", data)

	for {
		select {
		case <-ctx.Done():
			log.Println("Main function is exiting.")
			return
		default:
			time.Sleep(8 * time.Second)
			log.Printf("Random data generation is in progress...\n")
		}
	}
}
