package server

import (
	"go_rest_api/pkg"
	"log"
	"net/http"
)

type Server struct {
	s  *http.Server
	us root.UserService
}

func NewServer(u root.UserService) *Server {
	return &Server{
		&http.Server{Addr: ":8080"},
		u,
	}
}

func (srv *Server) Start() {
	log.Println("Listening on port 8080")
	log.Fatal("http.ListenAndServe: ", srv.s.ListenAndServe())
}
