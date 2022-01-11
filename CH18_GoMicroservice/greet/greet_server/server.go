package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	// package name "import path"
	greet "github.com/example/greet/greetpb" // specify package name - package name and declaration is not the same
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type Server struct {
}

// Implement Greet Service

// 1. Greet endpoint
func (s *Server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	log.Println("Received a Greet request :", req)

	greeting := req.GetGreeting()
	response := &greet.GreetResponse{
		Result: fmt.Sprintf("Hello, %s %s", greeting.GetFirstName(), greeting.GetLastName()),
	}

	return response, nil
}

// 2. Greet Many Times endpoint
func (s *Server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	log.Println("Receive a Greet Many Times request :", req)

	greeting := req.GetGreeting()
	n := req.GetTimes()

	var i int32
	for i = 0; i < n; i++ {
		response := &greet.GreetManyTimesResponse{
			Result: fmt.Sprintf("%s - Hello, %s %s", strconv.Itoa(int(i)), greeting.GetFirstName(), greeting.GetLastName()),
		}
		stream.Send(response)

		time.Sleep(time.Millisecond * 1000) // Sleep for just 1 sec
	}

	return nil
}

// 3. Long Greet endpoint
func (s *Server) LongGreet(stream greet.GreetService_LongGreetServer) error {
	log.Println("Receive a Long Greet stream")

	c := make(chan string)

	go func() {
		var b bytes.Buffer

		b.WriteString(fmt.Sprintf("Hello,\n"))
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occur while reading streaming : %v", err)
			}
			greeting := msg.GetGreeting()
			b.WriteString(fmt.Sprintf("\t%s %s\n", greeting.GetFirstName(), greeting.GetLastName()))

			time.Sleep(time.Millisecond * 1000) // Sleep for just 1 sec
		}

		c <- b.String()
	}()

	result := <-c
	return stream.SendAndClose(&greet.LongGreetResponse{
		Result: result,
	})
}

var wg sync.WaitGroup

// 4. Greet Everyone endpoint
func (s *Server) GreetEveryone(stream greet.GreetService_GreetEveryoneServer) error {
	log.Println("Receive a Greet Everyone stream")

	// Buffered channel with capacity
	c := make(chan string, 5)

	// 1. Receive a bunch of message from client stream
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occur while reading client streaming : %v", err)
			}

			greeting := req.GetGreeting()
			c <- fmt.Sprintf("Hello, %s %s", greeting.GetFirstName(), greeting.GetLastName())
		}

		close(c)
	}()

	wg.Add(1)

	// 2. Send a bunch of message to server stream
	go func() {
		log.Printf("Processing %d data stream from the client\n", cap(c))

		time.Sleep(5 * time.Second) // Wait for 5 seconds, until channel are full

		log.Println("Ready to send ...")
		for {
			if result, ok := <-c; ok {
				err := stream.Send(&greet.GreetEveryoneResponse{
					Result: result,
				})
				if err != nil {
					log.Fatalf("Error occur while sending client streaming : %v", err)
				}
			} else {
				// Channel is closed
				break
			}
		}

		wg.Done()
	}()

	wg.Wait()
	log.Println("Done send requests to the Client ...")

	return nil
}

// 5. Greet with Deadline endpoint
func (s *Server) GreetDeadline(ctx context.Context, req *greet.GreetDeadlineRequest) (*greet.GreetDeadlineResponse, error) {
	log.Println("Received a Greet request :", req)

	// Perform a heavy task
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		// Check client deadline if already exceeded
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Terminate the process. Deadline already exceeded")
			return nil, status.Errorf(codes.DeadlineExceeded, "Terminate the process")
		}
	}

	greeting := req.GetGreeting()
	response := &greet.GreetDeadlineResponse{
		Result: fmt.Sprintf("Hello, %s %s", greeting.GetFirstName(), greeting.GetLastName()),
	}

	return response, nil
}

var sslEnable = true

func main() {
	log.Println("Running gRPC server ... at port 50851")
	// Create a TCP connection
	lis, err := net.Listen("tcp", "127.0.0.1:50851")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	opts := []grpc.ServerOption{}
	if sslEnable {
		// Setting TLS in Server; using server.crt - server certificate and server.pem - server key (public and private)
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pkcs8.pem"

		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed to read server certificates: %v", sslErr)
		}

		opts = append(opts, grpc.Creds(creds))
	}
	// Create a new gRPC Server
	s := grpc.NewServer(opts...)
	// Register service into the Server
	greet.RegisterGreetServiceServer(s, &Server{})

	// Register service ...

	// Bind a TCP connection into gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}

}
