package server

import (
	"net/http"

	"github.com/VicFlores/src/middlewares"
	"github.com/VicFlores/src/router"
)

type Server struct {
	Port   string
	router *router.Router
}

// Definimos un puerto en el struct Server
func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		router: router.NewRouter(),
	}
}

func (s *Server) Handle(path, method string, handler http.HandlerFunc) {
	_, exist := s.router.Rules[path]

	if !exist {
		s.router.Rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.Rules[path][method] = handler
}

func (s *Server) AddMiddleware(hf http.HandlerFunc, middlewares ...middlewares.Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		hf = m(hf)
	}

	return hf
}

// Hacemos que el server este atento a las peticiones
func (s *Server) Listen() error {
	http.Handle("/", s.router)

	err := http.ListenAndServe(s.Port, nil)

	if err != nil {
		return err
	}

	return nil
}
