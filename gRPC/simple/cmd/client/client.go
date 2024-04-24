package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"simple/proto/message"
	"simple/proto/service"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := service.NewApiServiceClient(conn)

	ctx := context.Background()
	request := message.StringMessage{
		Value: "Hello ",
	}

	stringMessage, err := client.Echo(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(stringMessage)
}
