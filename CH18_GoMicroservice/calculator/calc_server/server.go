package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"sync"

	calc "github.com/example/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = 8081
)

type Server struct {
}

// Implement Calculator Service

// 1. Sum endpoint
func (s *Server) Sum(ctx context.Context, req *calc.SumRequest) (*calc.SumResponse, error) {
	log.Println("Receive a Sum request", req)

	numbers := req.GetNumbers()
	return &calc.SumResponse{
		Result: numbers.GetFirst() + numbers.GetSecond(),
	}, nil
}

// 2. Prime decompose endpoint
func (s *Server) Decompose(req *calc.DecomposeRequest, stream calc.CalcService_DecomposeServer) error {
	log.Println("Receive a Decompose request", req)

	n := req.GetNumber()
	var k int64 = 2
	for n > 1 {
		if n%k == 0 {
			res := &calc.DecomposeResponse{
				PrimeNumber: k,
			}

			stream.Send(res)

			n = n / k
		} else {
			k++
		}
	}

	return nil
}

// 3. Average calculation endpoint
func (s *Server) Average(stream calc.CalcService_AverageServer) error {

	c := make(chan float64)

	go func() {
		count, total := 0, 0
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading streaming : %v", err)
			}
			total += int(msg.GetNumber())
			count++
		}

		c <- float64(total) / float64(count)
	}()

	result := <-c
	return stream.SendAndClose(&calc.AverageResponse{
		Result: result,
	})
}

var wg sync.WaitGroup

// 4. Find Max endpoint
func (s *Server) FindMax(stream calc.CalcService_FindMaxServer) error {

	c := make(chan int32)
	// 1. Receive a bunch of numbers from client stream
	go func() {
		max := int32(0)
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading streaming : %v", err)
			}
			number := req.GetNumber()
			if number > max {
				c <- number
				max = number
			}
		}

		close(c)
	}()

	wg.Add(1)

	// 2. Send a bunch of max numbers to server stream
	go func() {
		for {
			if max, ok := <-c; ok {
				stream.Send(&calc.FindMaxResponse{
					MaxNumber: max,
				})
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

func (s *Server) SquareRoot(ctx context.Context, req *calc.SquareRootRequest) (*calc.SquareRootResponse, error) {
	number := req.GetNumber()
	if number < 0 {
		errMsg := fmt.Sprintf("Invalid number. %d is negative", number)
		return nil, status.Errorf(codes.InvalidArgument, errMsg)
	}

	return &calc.SquareRootResponse{
		SquareNumber: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	log.Println("Running gRPC server ... at port", port)

	lis, err := net.Listen("tcp", fmt.Sprint("127.0.0.1:", port))
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	// Register a service implementation ...
	calc.RegisterCalcServiceServer(s, &Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	log.Println("Server ready to serve ...")
	if err := s.Serve(lis); err != nil { // Wait for client request
		log.Fatalf("Failed to serve : %v", err)
	}
}
