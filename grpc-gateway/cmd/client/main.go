package main

import (
	"context"
	"gateway/protogen/golang/users"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {

	serviceAddr := ":50051"
	conn, err := grpc.Dial(serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	mux := runtime.NewServeMux()
	if err = users.RegisterUsersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("could not register handler: %v", err)
	}

	addr := ":8080"
	log.Printf("listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("could not listen and serve: %v", err)
	}

}
