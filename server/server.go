package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkung/srv/config"
)

type Server struct {
	server *http.Server
	Engine *gin.Engine
}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.server.Close()
}

func NewServer(customizer ServerCustomizer) *Server {
	cfg := config.New()
	router := gin.Default()
	customizer.Register(router)

	router.Run()

	s := &Server{
		server: &http.Server{
			Addr:    cfg.ServerAddr,
			Handler: router,
		},
		Engine: router,
	}

	return s
}
