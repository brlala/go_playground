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

	// CRUD API
	blogId := CreateBlog(c)
	ReadBlog(c, "asncas")
	ReadBlog(c, "607b1395fbf3c071a83cc17e")
	ReadBlog(c, blogId)

	// Update Blog
	newBlog := &blogpb.Blog{
		//Id: blogId,
		Id:       blogId,
		AuthorId: "Changed Author 1",
		Content:  "Edited blog content",
		Title:    "Content of the blog with addition",
	}
	UpdateBlog(c, newBlog)

	// Delete Blog
	DeleteBlog(c, blogId)
}

func CreateBlog(c blogpb.BlogServiceClient) string {
	// Create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{AuthorId: "Stephan", Title: "First Blog", Content: "Content of first blog"}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	return createBlogRes.GetBlog().GetId()
}
func ReadBlog(c blogpb.BlogServiceClient, blogId string) {
	// Read Blog
	fmt.Println("Reading the blog")
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogId})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
		return
	}
	fmt.Printf("Blog was read: %v\n", res)
}

func UpdateBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog) {
	// Update Blog
	fmt.Println("Updating the blog")
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		fmt.Printf("Error happened while updating: %v\n", err)
		return
	}
	fmt.Printf("Blog was updated: %v\n", res)
}

func DeleteBlog(c blogpb.BlogServiceClient, blogId string) {
	// Update Blog
	fmt.Println("Deleting the blog")
	res, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})
	if err != nil {
		fmt.Printf("Error happened while deleting: %v\n", err)
		return
	}
	fmt.Printf("Blog was deleted: %v\n", res)
}
