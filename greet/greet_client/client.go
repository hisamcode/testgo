package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hisamcode/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("hello i'am a client");

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure());

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)
	doServerStreaming(c)
	
}

func doUnary(c greetpb.GreetServiceClient)  {
	fmt.Println("Starting to do a unary RPC....")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hisam",
			LastName: "Maulana",
		},
	}
	
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC....")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hisam",
			LastName: "Maulana",
		},
	}
	// response stream
	rs, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
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
		log.Printf("Response from GreatManyTimes: %v", msg.GetResult())
	}
}