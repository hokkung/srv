package server

import "github.com/gin-gonic/gin"

type ServerCustomizer interface {
	Register(engine *gin.Engine)
}
