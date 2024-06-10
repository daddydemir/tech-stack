package main

import (
	"gateway/internal"
	"gateway/protogen/golang/users"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	const addr = ":50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	db := internal.NewDB()
	service := internal.NewUserService(db)

	users.RegisterUsersServer(server, &service)

	log.Printf("server listening at %s", addr)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
