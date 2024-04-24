package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"simple/proto/service"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	service.RegisterApiServiceServer(server, ApiServer{})
	reflection.Register(server)

	server.Serve(listen)
}
