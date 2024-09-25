package main

import (
	"fmt"
	"go-test/common"
	"go-test/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/users", func(context *gin.Context) {
		limitStr := context.DefaultQuery("limit", "10")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, users.FindUsersResponse{
				Error: common.ToPtr("Invalid 'limit' query parameter"),
			})
			return
		}

		mockUsers := users.GenerateMockUsers(limit)
		context.JSON(http.StatusOK, users.FindUsersResponse{
			Data:  &mockUsers,
			Total: common.ToPtr(len(mockUsers)),
		})
	})

	err := engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
