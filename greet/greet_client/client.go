package main

import (
	"context"
	"fmt"
	"github.com/parsarsm/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello! I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting doUnary!")
	req := greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Parsa",
		LastName:  "Rostami",
	}}
	response, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %+v", err)
	}
	log.Printf("Response from Greet : %+v\n", response)
}
