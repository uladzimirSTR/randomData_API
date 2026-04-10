package randomdata

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/uladzimirSTR/randomData_API/dbase"
	wp "github.com/uladzimirSTR/randomData_API/workerPool"
)

func RandomDataUsers(
	pool *pgxpool.Pool,
	firstRun bool,
) {
	// Initialize a worker pool with a specified number of workers
	var wpool wp.Pool = wp.New(5)
	wpool.Make(5)

	if firstRun {
		log.Println("Creating 'users' table...")
		db.CreateTable(pool, "users", "random_data", []db.Column{
			{Name: "id", Type: "BIGSERIAL", NotNull: true},
			{Name: "email", Type: "TEXT", NotNull: true},
			{Name: "name", Type: "TEXT"},
			{Name: "created_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
			{Name: "updated_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
		}, []string{"id"})
	}

	for _ = range [10]struct{}{} {
		wpool.Handle(func() {
			users := GenerateRandomUsers(rand.Intn(100))

			if len(users) == 0 {
				users = GenerateRandomUsers(1)
			}

			rows := make([]string, 0, len(users))

			for _, user := range users {
				row := fmt.Sprintf(
					"%d, '%s', '%s', '%s', '%s'",
					user.ID,
					user.Email,
					user.Name,
					user.CreatedAt,
					user.UpdatedAt,
				)

				rows = append(rows, row)
			}

			db.InsertValues(
				pool,
				"random_data",
				"users",
				[]string{"id", "email", "name", "created_at", "updated_at"},
				[]string{"id"},
				rows,
			)
		})
	}

	wpool.Wait()

}
