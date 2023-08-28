package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkung/srv/server"
)

type Custom struct {
}

func (custom *Custom) Register(s *server.Server) {
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func main() {
	cus := &Custom{}
	s := server.NewServer(cus)
	err := s.Start()
	if err != nil {
		fmt.Println(err)
	}
}
