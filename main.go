package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	db "github.com/uladzimirSTR/randomData_API/dbase"
	rnd "github.com/uladzimirSTR/randomData_API/randomData"
	wp "github.com/uladzimirSTR/randomData_API/workerPool"
)

const TIME_WORK = 5 * time.Minute

func randomData(ctx context.Context) {

	var wpool wp.Pool = wp.New(5)
	wpool.Make(5)

	db.CreateTable(ctx, "users", "random_data", []db.Column{
		{Name: "id", Type: "BIGSERIAL", NotNull: true},
		{Name: "email", Type: "TEXT", NotNull: true},
		{Name: "name", Type: "TEXT"},
		{Name: "created_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
		{Name: "updated_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
	}, []string{"id"})

l:
	for {
		select {
		case <-ctx.Done():
			log.Println("Random data generation stopped.")
			break l
		default:
			wpool.Handle(func() {
				users := rnd.GenerateRandomUsers(rand.Intn(100))

				if len(users) == 0 {
					users = rnd.GenerateRandomUsers(1)
				}

				// log.Printf("line: %+v\n", users)
				db.InsertUsers(ctx, "random_data", "users", users)
			})
		}
	}

	wpool.Wait()

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIME_WORK)
	defer cancel()

	go randomData(ctx)

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
