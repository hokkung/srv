package server

import (
	"fmt"
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

func (s *Server) Stop() {
	err := s.server.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func NewServer(customizer ServerCustomizer) (*Server, error) {
	cfg := config.New()
	router := gin.Default()
	server := &Server{
		server: &http.Server{
			Addr:    cfg.ServerAddr,
			Handler: router,
		},
		Engine: router,
	}

	customizer.Register(server)

	return server, nil
}

func ProvideServer(customizer ServerCustomizer) (*Server, func(), error) {
	srv, err := NewServer(customizer)
	if err != nil {
		return nil, func() {}, err
	}
	return srv, srv.Stop, nil
}
