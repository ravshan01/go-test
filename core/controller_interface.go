package core

import "github.com/gin-gonic/gin"

type IController interface {
	init(engine *gin.Engine)
}
