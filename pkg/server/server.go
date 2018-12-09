package server

import (
	"go_rest_api/pkg"
	"log"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(u root.UserService) *Server {
	NewUserRouter(u, handleGroup("/u"))
	return &Server{&http.Server{Addr: ":8080"}}
}

func (srv *Server) Start() {
	log.Println("Listening on port 8080")
	log.Fatal("http.ListenAndServe: ", srv.ListenAndServe())
}

func handleGroup(path string) func(string, http.Handler) {
	return func(subpath string, handler http.Handler) {
		http.Handle(path+subpath, handler)
	}
}
