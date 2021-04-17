package main

import (
	"context"
	"fmt"
	"go_playground/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	tls := false
	// grpc by default has SSL, this line is to remove it, disable SSL and open a connection to the port
	opts := grpc.WithInsecure()

	// using SSL file
	if tls {
		certFile := "ssl/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer keyword defer this code until the very end of the code
	defer cc.Close()

	// create a new client
	c := blogpb.NewBlogServiceClient(cc)
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{AuthorId: "Stephan", Title: "First Blog", Content: "Content of first blog"}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
}
