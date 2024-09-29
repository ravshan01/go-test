package core

import "github.com/gin-gonic/gin"

type IController interface {
	Init(engine *gin.Engine)
}
