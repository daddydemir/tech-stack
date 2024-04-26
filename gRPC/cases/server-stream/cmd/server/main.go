package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"server-stream/proto/service"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	service.RegisterServerServiceServer(server, Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
