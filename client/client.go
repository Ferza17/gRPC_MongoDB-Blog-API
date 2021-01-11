package main

import (
	"context"
	"fmt"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello i'm a Blog client")
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := blog_proto.NewBlogServiceClient(cc)
	doUnary(c)
}

func doUnary(c blog_proto.BlogServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blog_proto.CreateBlogRequest{
		Blog: &blog_proto.Blog{
			AuthorId: "Fery Reza",
			Title: "This is Title 1",
			Content: "Content of the Blog",
		},
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Blog RPC: %v", err)
	}

	log.Printf("Response From Greet: %v", res.GetBlog())
}
