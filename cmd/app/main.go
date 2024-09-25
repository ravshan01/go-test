package main

import (
	"fmt"
	"go-test/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", func(c *gin.Context) {
		mockUsers := users.GenerateMockUsers(10)
		c.JSON(http.StatusOK, mockUsers)
	})

	err := server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
