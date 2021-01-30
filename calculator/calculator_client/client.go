package main

import (
	"context"
	"fmt"
	"github.com/parsarsm/go-grpc/calculator/calculatorpb"
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

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting sum doUnary!")
	req := calculatorpb.SumRequest{
		FirstNumber:  13,
		SecondNumber: 17,
	}
	response, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %+v", err)
	}
	log.Printf("Response from Sum : %+v\n", response)
}
