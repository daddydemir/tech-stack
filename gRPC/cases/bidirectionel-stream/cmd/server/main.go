package main

import (
	"bidirectional-stream/proto/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service.RegisterBidirectionalStreamServer(server, Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
