package main

import (
	"context"
	"simple/proto/message"
	"simple/proto/service"
)

type ApiServer struct {
	service.UnimplementedApiServiceServer
}

func (ApiServer) Echo(ctx context.Context, m *message.StringMessage) (*message.StringMessage, error) {

	stringMessage := message.StringMessage{Value: m.Value + " world!"}
	return &stringMessage, nil
}
