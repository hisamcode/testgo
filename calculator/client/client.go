package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hisamcode/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("hello i'am a client");

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure());

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	
}

func doUnary(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting to do a unary RPC....")
	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			Number_1: 3,
			Number_2: 10,
		},
	}
	
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Response from sum: %v", res.Result)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting to do a doPrimeNumberDecomposition server streaming RPC....")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 6000000,
	}
	// response stream
	rs, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}
	for {
		msg, err := rs.Recv()
		if err == io.EOF {
			// we 've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreatManyTimes: %v", msg.GetPrimeFactor())
	}
}