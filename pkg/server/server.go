package server

import (
	"go_rest_api/pkg"
	"log"
	"net/http"
)

type Server struct {
	s *http.Server
}

func NewServer(u root.UserService) *Server {
	NewUserRouter(u, group("/u"))
	return &Server{
		&http.Server{Addr: ":8080"},
	}
}

func (srv *Server) Start() {
	log.Println("Listening on port 8080")
	log.Fatal("http.ListenAndServe: ", srv.s.ListenAndServe())
}

func group(path string) func(string, http.Handler) {
	return func(subpath string, handler http.Handler) {
		http.Handle(path+subpath, handler)
	}
}
