package users

import (
	"github.com/gin-gonic/gin"
	"go-test/core"
	"strconv"
)

type UsersController struct {
	usersService *UsersService
}

func NewUsersController() *UsersController {
	return &UsersController{
		usersService: &UsersService{},
	}
}

func (c *UsersController) init(engine *gin.Engine) {
	engine.GET("/users", func(ctx *gin.Context) {
		limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
		if err != nil {
			ctx.JSON(400, core.BadRequestResponse{core.Response{
				Error: core.ResponseError{Message: "Invalid limit"},
			}})
			return
		}

		offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
		if err != nil {
			ctx.JSON(400, core.BadRequestResponse{core.Response{
				Error: core.ResponseError{Message: "Invalid offset"},
			}})
			return
		}

		params := UsersServiceFindParams{
			Limit:  limit,
			Offset: offset,
		}

		users, err := c.usersService.find(params)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, core.Response{Data: users})
	})

	//engine.GET("/users/:id", c.findById)
	//engine.POST("/users", c.create)
	//engine.PUT("/users", c.update)
	//engine.PATCH("/users/:id", c.partialUpdate)
	//engine.DELETE("/users/:id", c.delete)
}
