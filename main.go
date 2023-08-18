package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkung/srv/server"
)

type Custom struct {
}

func (custom *Custom) Register(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func main() {
	cus := &Custom{}
	s := server.NewServer(cus)
	s.Start()
}
