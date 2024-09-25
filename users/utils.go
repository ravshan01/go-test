package users

import (
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"math/rand"
)

func GenerateMockUsers(n int) []User {
	fake := faker.New()
	var users []User

	for i := 0; i < n; i++ {
		user := User{
			Id:    uuid.New().String(),
			Name:  fake.Person().Name(),
			Email: fake.Internet().Email(),
			Age:   rand.Intn(100),
		}

		users = append(users, user)
	}

	return users
}
