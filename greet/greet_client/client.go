package main

import (
	"context"
	"fmt"
	"go_playground/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	// grpc by default has SSL, this line is to remove it, disable SSL and open a connection to the port
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer keyword defer this code until the very end of the code
	defer cc.Close()

	// create a new client
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)
	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDirectionalStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephanie",
			LastName:  "Hawking",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.GetResult())
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephan",
			LastName:  "Hawking"}}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		} else if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from Greet: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}

	request := []*greetpb.LongGreetRequest{
		{Greeting: &greetpb.Greeting{FirstName: "Stephane", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "John", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Lucy", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Mark", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Piper", LastName: "Row"}},
	}

	// We iterate over our slice and send each message individually
	for _, req := range request {
		fmt.Printf("Sending req: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("error while sending request to LongGreet RPC: %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBiDirectionalStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bi Directional Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{Greeting: &greetpb.Greeting{FirstName: "Stephane", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "John", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Lucy", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Mark", LastName: "Row"}},
		{Greeting: &greetpb.Greeting{FirstName: "Piper", LastName: "Row"}},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// we recieve a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()
	// block until everything is done(channel is closed)
	<-waitc
}
