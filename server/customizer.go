package server

type ServerCustomizer interface {
	Register(engine *Server)
}
