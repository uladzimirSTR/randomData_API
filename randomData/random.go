package randomdata

import (
	"log"
	"math/rand"
	"time"

	obj "github.com/uladzimirSTR/randomData_API/objects"
)

var currentTime = time.Now()

func GenerateRandomUser() obj.User {
	name := names[rand.Intn(len(names))]

	return obj.User{
		ID:        rand.Intn(1_000_000) + rand.Intn(1_000_000) + rand.Intn(1_000_000),
		Email:     name + "@" + domains[rand.Intn(len(domains))],
		Name:      name,
		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}
}

func GenerateRandomUsers(count int) []obj.User {
	users := make([]obj.User, count)
	for i := 0; i < count; i++ {
		users[i] = GenerateRandomUser()
	}

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // Simulate some processing time
	log.Printf("Generated %d random users\n", count)
	return users
}
