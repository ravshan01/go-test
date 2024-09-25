package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.GET("/users", func(c *gin.Context) {
		users := generateMockUsers(10)
		c.JSON(http.StatusOK, users)
	})

	err := server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func generateMockUsers(n int) []User {
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
