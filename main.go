package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/core"
	"go-test/users"
	"net/http"
)

func main() {
	err := bootstrap()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func bootstrap() error {
	return startServer("localhost", "8080")
}

func startServer(host, port string) error {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var usersController core.IController = users.NewUsersController()
	usersController.Init(engine)

	return engine.Run(fmt.Sprintf("%s:%s", host, port))
}
