package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/config"
	"go-test/core"
	"go-test/users"
	"net/http"
)

// parse
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = bootstrap(cfg)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func bootstrap(cfg *config.Config) error {
	err := startServer(cfg.Host, cfg.Port)
	return err
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
