package main

import (
	"go_rest_api/internal/bolt"
	"go_rest_api/internal/crypto"
	"go_rest_api/internal/server"
	"go_rest_api/internal/sessionStore"
	"log"
)

const (
	USERS = "users"
)

func main() {
	blt, err := bolt.NewDatabase()
	if err != nil {
		log.Fatalln("unable to connect to boltdb")
	}
	defer blt.Close()

	h := crypto.Hash{}
	userService := bolt.NewUserService(blt, USERS, &h)
	store := sessionStore.NewStore(USERS)
	userStore := sessionStore.NewUserStore(store)
	s := server.NewServer(userService, userStore)

	s.Start()
}
