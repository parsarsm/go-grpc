package main

import (
	"context"
	"fmt"
	"github.com/parsarsm/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet func was just invoked with %+v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	response := greetpb.GreetResponse{Result: result}
	return &response, nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
