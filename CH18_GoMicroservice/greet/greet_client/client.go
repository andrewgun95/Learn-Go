package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	greet "github.com/example/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var wg sync.WaitGroup

var sslEnable = true

func main() {
	log.Println("Try connecting to gRPC server ... at port 50851")

	opts := grpc.WithInsecure()
	if sslEnable {
		// Get CA (Certificate Authority) trusted certificate
		certFile := "ssl/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Failed to read trusted certificates: %v", sslErr)
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	// Dialing to gRPC Server connection
	cli_conn, err := grpc.Dial("127.0.0.1:50851", opts)
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}

	// Close the connection after exited the program
	defer cli_conn.Close()

	// Get service from the connection
	c := greet.NewGreetServiceClient(cli_conn)

	doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)

	// Perform RPC calls with deadline
	// 1. Success
	// doUnaryWithDeadline(c, 5)
	// 2. Fail
	// doUnaryWithDeadline(c, 1)
}

func doUnary(c greet.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC Call ...")

	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "Andrew",
			LastName:  "Gun",
		},
	}

	log.Println("Sending a Greet request : ", req)
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}
	log.Printf("Got response Greet : %v", res.Result)
}

func doServerStreaming(c greet.GreetServiceClient) {
	log.Println("Starting to do a Server Streaming RPC Call ...")

	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: "Andrew",
			LastName:  "Gun",
		},
		Times: 10,
	}

	log.Println("Sending a Greet Many Times request : ", req)
	clientStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	wg.Add(1)

	go func() {
		for {
			msg, err := clientStream.Recv()
			if err == io.EOF { // Reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading stream : %v", err)
			}
			log.Printf("Got response Greet Many Times : %v", msg.GetResult())
		}

		wg.Done()
	}()

	wg.Wait()
}

func doClientStreaming(c greet.GreetServiceClient) {
	log.Println("Starting to do a Client Streaming RPC Call ...")

	clientStream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	requests := []*greet.LongGreetRequest{
		{
			Greeting: &greet.Greeting{
				FirstName: "Andrew",
				LastName:  "Gun",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "James",
				LastName:  "Bond",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Miss",
				LastName:  "Moneypenny",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Peter",
				LastName:  "Robert",
			},
		},
	}

	for _, req := range requests {
		log.Println("Sending a Long Greet request : ", req)
		clientStream.Send(req)
		time.Sleep(time.Millisecond * 100) // Sleep just a little bit 0.1 secs
	}

	log.Println("Wait until got response from the Server ...")
	res, _ := clientStream.CloseAndRecv() // Wait until got response

	log.Printf("Got response Long Greet : %v", res.GetResult())
}

func doBiDiStreaming(c greet.GreetServiceClient) {
	log.Println("Starting to do a BiDi Streaming RPC Call ...")

	// 1. Create a client stream by invoking the client service
	clientStream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	requests := []*greet.GreetEveryoneRequest{
		{
			Greeting: &greet.Greeting{
				FirstName: "Andrew",
				LastName:  "Gun",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "James",
				LastName:  "Bond",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Miss",
				LastName:  "Moneypenny",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Peter",
				LastName:  "Robert",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Jacob",
				LastName:  "Smith",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Alice",
				LastName:  "Key",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Marcell",
				LastName:  "Key",
			},
		},
	}

	wg.Add(1)

	// 2. Send a bunch of messages to a client stream
	go func() {
		for _, req := range requests {

			time.Sleep(1 * time.Second) // Sleep just a little bit

			log.Println("Sending a Greet Everyone request : ", req)
			err := clientStream.Send(req)
			if err != nil {
				log.Fatalf("Error occur while sending server streaming : %v", err)
			}
		}

		clientStream.CloseSend()
	}()

	// 3. Receive a bunch of message from server stream
	go func() {
		for {
			res, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occur while reading server streaming : %v", err)
			}

			log.Printf("Got response Greet Everyone : %v", res.GetResult())
		}
		wg.Done()
	}()

	log.Println("Wait until got response from the Server ...")
	wg.Wait()
}

func doUnaryWithDeadline(c greet.GreetServiceClient, timeout int) {
	log.Println("Starting to do a Unary with Deadline RPC Call ...")

	req := &greet.GreetDeadlineRequest{
		Greeting: &greet.Greeting{
			FirstName: "Andrew",
			LastName:  "Gun",
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	// Cancel the request - at the End
	defer cancel()

	log.Println("Sending a Greet request : ", req)

	res, err := c.GreetDeadline(ctx, req)
	if err != nil {
		if stat, ok := status.FromError(err); ok {
			// User-defined error
			if stat.Code() == codes.DeadlineExceeded {
				// Handle error timeout
				log.Printf("Request timeout. Can't got the response : %v", stat.Message())
			} else {
				log.Printf("Unexpected error : %v\n", stat.Message())
			}
		} else {
			// System error
			log.Fatalf("Error occurs : %v", err)
		}

		return // Error found, no response
	}

	log.Printf("Got response Greet : %v", res.Result)
}
