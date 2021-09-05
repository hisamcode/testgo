package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/hisamcode/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Sum function was invokeed with %v\n", req)
	number1 := req.Calculator.GetNumber_1()
	number2 := req.Calculator.GetNumber_2()
	sum := number1 + number2

	result := "sum of " + strconv.Itoa(int(number1)) + " + " + strconv.Itoa(int(number2)) + " is " + strconv.Itoa(int(sum))
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invokeed with %v\n", req)
	number := req.GetNumber()
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
			// sleep 1 detik setiap eksekusi
			time.Sleep(1000 * time.Millisecond)
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func main() {

	//The service at the port 8000 is the UI only. If you wanna use gRPC you have to call it in the port 9080. Never 8080, 8000, 6080, and so on.
	//https://stackoverflow.com/questions/63372312/rpc-error-code-unavailable-desc-connection-closed-in-go-code
	httpPort := "9080"

	lis, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("test dari app")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
