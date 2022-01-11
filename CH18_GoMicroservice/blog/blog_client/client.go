package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	blog "github.com/example/blog/blogpb"
	"google.golang.org/grpc"
)

const (
	port = "8081"
)

var wg sync.WaitGroup

func main() {
	log.Println("Try connecting to gRPC server ... at port", port)

	opts := grpc.WithInsecure()
	// Dialing to gRPC Server connection
	cli_conn, err := grpc.Dial(fmt.Sprint("127.0.0.1:", port), opts)
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}

	// Close the connection after exited the program
	defer cli_conn.Close()

	// Get service from the connection
	c := blog.NewBlogServiceClient(cli_conn)

	createReq := &blog.CreateBlogRequest{
		Blog: &blog.Blog{
			AuthorId: "xyz",
			Title:    "first book",
			Content:  "this is best book ever",
		},
	}

	log.Println("Send a Blog :", createReq)
	createRes, err := c.CreateBlog(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}
	log.Println("Got response Blog :", createRes)

	blogId := createRes.GetResult().GetId()

	// Return a response
	readRequest := &blog.ReadBlogRequest{
		BlogId: blogId,
	}

	log.Println("Send a Blog :", readRequest)
	readRes, err := c.ReadBlog(context.Background(), readRequest)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}
	log.Println("Got response Blog :", readRes)

	// Return an error - Not Found
	// if _, err := c.ReadBlog(context.Background(), &blog.ReadBlogRequest{BlogId: "60a34e173067add64535d5e1"}); err != nil {
	// 	log.Fatalf("Error occurs : %v", err)
	// }

	time.Sleep(3 * time.Second)

	updateReq := &blog.UpdateBlogRequest{
		Blog: &blog.Blog{
			Id:       blogId,
			AuthorId: "xyz,abc (edited)",
			Title:    "first book (edited)",
			Content:  "absolutely the best book ever (edited)",
		},
	}
	log.Println("Send a Blog :", updateReq)
	updateRes, err := c.UpdateBlog(context.Background(), updateReq)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}
	log.Println("Go response Blog :", updateRes)

	time.Sleep(3 * time.Second)

	deleteReq := &blog.DeleteBlogRequest{
		BlogId: blogId,
	}
	log.Println("Send a Blog :", deleteReq)
	deleteRes, err := c.DeleteBlog(context.Background(), deleteReq)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}
	log.Println("Go response Blog :", deleteRes)

	listReq := &blog.ListBlogRequest{}
	log.Println("Send a Blog :", listReq)
	clientStream, err := c.ListBlog(context.Background(), listReq)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	wg.Add(1)

	go func() {
		for {
			listRes, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading stream : %v", err)
			}

			log.Println("Got response Blog :", listRes)
		}

		wg.Done()
	}()

	wg.Wait() // Wait for 1
}
