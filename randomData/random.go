package randomdata

import (
	"log"
	"math/rand"
	"time"
)

var currentTime = time.Now()

func GenerateRandomUser() User {
	name := names[rand.Intn(len(names))]

	return User{
		ID:        rand.Intn(1_000_000) + rand.Intn(1_000_000) + rand.Intn(1_000_000),
		Email:     name + "@" + domains[rand.Intn(len(domains))],
		Name:      name,
		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}
}

func GenerateRandomUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = GenerateRandomUser()
	}

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // Simulate some processing time
	log.Printf("Generated %d random users\n", count)
	return users
}
