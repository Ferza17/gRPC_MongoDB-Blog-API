package main

import (
	"context"
	"fmt"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello i'm a Blog client")
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	c := blog_proto.NewBlogServiceClient(cc)
	//doUnary(c)
	//doUnarySearchById(c)
	//doUnaryUpdate(c)
	//doDeleteBlog(c)
	listBlog(c)

	_ = cc.Close()
}

func doUnary(c blog_proto.BlogServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blog_proto.CreateBlogRequest{
		Blog: &blog_proto.Blog{
			AuthorId: "Aditya",
			Title:    "My Awesome Blog",
			Content:  "Content of the Blog",
		},
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Blog RPC: %v", err)
	}

	log.Printf("Response From Greet: %v", res.GetBlog())
}

func doUnarySearchById(c blog_proto.BlogServiceClient) {
	fmt.Println("About to start doUnarySearchById...")
	req := &blog_proto.ReadBlogRequest{BlogId: "5ffd70d6f0a3836bbc1dafdc"}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While ReadBlog!")
	}

	log.Println("Response From ReadBlog : ", res)
}

func doUnaryUpdate(c blog_proto.BlogServiceClient) {
	newBlog := &blog_proto.Blog{
		Id:       "5ffd3775164fb066d0d38eb7",
		AuthorId: "Changed Author 3",
		Title:    "Changed Title",
	}

	res, err := c.UpdateBlog(context.Background(), &blog_proto.UpdateBlogRequest{
		Blog: newBlog,
	})
	if err != nil {
		log.Fatalf("Error While UpdateBlog!")
	}

	log.Println("Response From ReadBlog : ", res)
}

func doDeleteBlog(c blog_proto.BlogServiceClient) {
	fmt.Println("Starting to do a doDeleteBlog RPC...")

	res, err := c.DeleteBlog(context.Background(), &blog_proto.DeleteBlogRequest{BlogId: "5ffd9267885ef5ed12cc5f9e"})
	if err != nil {
		log.Println("error while DeleteBlog : ", err)
	}

	log.Println("Result :", res)
}

func listBlog(c blog_proto.BlogServiceClient) {
	fmt.Println("Starting to do a listBlog RPC...")

	stream, err := c.ListBlog(context.Background(), &blog_proto.ListBlogRequest{})
	if err != nil {
		log.Println("Error while Read Blog")
	}

	fmt.Sprintln(stream)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error While ListBlog")
		}

		fmt.Sprintln("Result : ", res.GetBlog())
	}
}
