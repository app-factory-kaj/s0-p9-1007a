package handlers

import (
	"context"

	"greeter/internal/gen"
)

const defaultName = "World"

type GreetingServer struct{}

func NewGreetingServer() *GreetingServer {
	return &GreetingServer{}
}

func (s *GreetingServer) GetGreeting(_ context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name
	if name == "" {
		name = defaultName
	}

	return gen.GetGreeting200JSONResponse{
		Name:    name,
		Message: "Hello, " + name + "!",
	}, nil
}
