package main

import (
	"context"
	"fmt"
	"go_playground/greet/greetpb"
	"google.golang.org/grpc"
	"log"
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
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephanie",
			LastName:  "Hawking",
		},
	}
	c.Greet(context.Background(), req)
}
