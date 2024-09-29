package users

import (
	"github.com/gin-gonic/gin"
	"go-test/core"
	"strconv"
)

type UsersController struct {
	usersService IUsersService
}

func NewUsersController() *UsersController {
	return &UsersController{
		usersService: &UsersService{},
	}
}

func (c *UsersController) Init(engine *gin.Engine) {
	engine.GET("/users", c.find)
	engine.GET("/users/:id", c.findById)
	engine.POST("/users", c.create)
	engine.PUT("/users", c.update)
	engine.PATCH("/users/:id", c.partialUpdate)
	engine.DELETE("/users/:id", c.delete)
}

func (c *UsersController) find(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(400, core.NewBadRequestResponse("Invalid limit"))
		return
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		ctx.JSON(400, core.NewBadRequestResponse("Invalid offset"))
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
	ctx.JSON(200, core.Response[[]User]{Data: &users})
}

func (c *UsersController) findById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.usersService.findById(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		ctx.JSON(404, core.NewBadRequestResponse("User not found"))
		return
	}
	ctx.JSON(200, core.Response[User]{Data: user})
}

func (c *UsersController) create(ctx *gin.Context) {
	var userCreate UserCreate
	err := ctx.ShouldBindJSON(&userCreate)
	if err != nil {
		ctx.JSON(400, core.NewBadRequestResponse(err.Error()))
		return
	}

	newUser, err := c.usersService.create(userCreate)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, core.Response[User]{Data: &newUser})
}

func (c *UsersController) update(ctx *gin.Context) {
	var user UserUpdate
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, core.NewBadRequestResponse(err.Error()))
		return
	}

	updatedUser, err := c.usersService.update(user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, core.Response[User]{Data: &updatedUser})
}

func (c *UsersController) partialUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var user UserPartialUpdate
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, core.NewBadRequestResponse(err.Error()))
		return
	}

	updatedUser, err := c.usersService.partialUpdate(id, user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, core.Response[User]{Data: &updatedUser})
}

func (c *UsersController) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.usersService.delete(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(204, nil)
}
