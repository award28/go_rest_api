package main

import (
	"go_rest_api/pkg/bolt"
	"go_rest_api/pkg/crypto"
	"go_rest_api/pkg/server"
	"log"
)

func main() {
	blt, err := bolt.NewDatabase()
	if err != nil {
		log.Fatalln("unable to connect to boltdb")
	}
	defer blt.Close()

	h := crypto.Hash{}
	u := bolt.NewUserService(blt, "users", &h)
	s := server.NewServer(u)

	s.Start()
}
