package main

import (
	"context"
	"fmt"
	"go_playground/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	// extracting info from greet request
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	// creating a response for it
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	// Get one response and send stream response
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello World")

	//port binding
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//create grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// binding the port to the grpc server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
