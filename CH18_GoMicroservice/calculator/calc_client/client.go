package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	calc "github.com/example/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = 8081
)

var wg sync.WaitGroup

func main() {
	cli_conn, err := grpc.Dial(fmt.Sprint("127.0.0.1:", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}
	defer cli_conn.Close()

	c := calc.NewCalcServiceClient(cli_conn)

	// exercise1(c)
	// exercise2(c)
	// exercise3(c)
	// exercise4(c)

	errorHandling(c)
}

func exercise1(c calc.CalcServiceClient) {
	req := &calc.SumRequest{}
	req.Numbers = &calc.Numbers{}
	req.Numbers.First = 12
	req.Numbers.Second = 13
	log.Println("Send a Sum", req.GetNumbers())

	res, _ := c.Sum(context.Background(), req)
	log.Println("Result of Sum is", res)
}

func exercise2(c calc.CalcServiceClient) {

	req := &calc.DecomposeRequest{
		Number: 80123120,
	}

	log.Println("Send a Decompose", req.GetNumber())

	clientStream, _ := c.Decompose(context.Background(), req)

	wg.Add(1)

	go func() {
		for {
			msg, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading stream : %v", err)
			}
			log.Printf("Got prime number : %v\n", msg.GetPrimeNumber())
		}

		wg.Done()
	}()

	wg.Wait()
}

func exercise3(c calc.CalcServiceClient) {
	clientStream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	requests := []*calc.AverageRequest{
		{Number: 1}, {Number: 2}, {Number: 3}, {Number: 4},
	}

	for _, req := range requests {
		log.Println("Send a Number", req.GetNumber())
		clientStream.Send(req)
	}

	res, _ := clientStream.CloseAndRecv()
	log.Println("Result of Average is", res.GetResult())
}

func exercise4(c calc.CalcServiceClient) {
	clientStream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("Error occurs : %v", err)
	}

	wg.Add(1)

	requests := []*calc.FindMaxRequest{
		{Number: 1}, {Number: 5}, {Number: 3}, {Number: 6}, {Number: 2}, {Number: 20},
	}

	// 1. Send a bunch of numbers to client stream
	go func() {
		for _, req := range requests {
			log.Println("Send a Number", req.GetNumber())
			err := clientStream.Send(req)
			if err != nil {
				log.Fatalf("Error occur while sending server streaming : %v", err)
			}
		}

		clientStream.CloseSend()
	}()

	// 2. Receive a bunch of numbers from server stream
	go func() {
		for {
			msg, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occurs while reading stream : %v", err)
			}

			log.Printf("Got max number : %v\n", msg.GetMaxNumber())
		}
	}()

	log.Println("Wait until got response from the Server ...")
	wg.Wait()
}

func errorHandling(c calc.CalcServiceClient) {

	// 1. Error call
	// result, err := squareRoot(-5, c)
	// if err != nil {
	// 	log.Fatalf("Error occurs : %v\n", err)
	// } else {
	// 	log.Printf("Got square number : %v\n", result)
	// }

	// 2. Success call
	result, err := squareRoot(5, c)
	if err != nil {
		log.Fatalf("Error occurs : %v\n", err)
	} else {
		log.Printf("Got square number : %v\n", result)
	}
}

func squareRoot(number int32, c calc.CalcServiceClient) (float64, error) {
	request := &calc.SquareRootRequest{
		Number: number,
	}

	res, err := c.SquareRoot(context.Background(), request)
	if err != nil {
		if stat, ok := status.FromError(err); ok {
			// User-defined error
			log.Println(stat.Code(), ":", stat.Message())
			if stat.Code() == codes.InvalidArgument {
				return 0, fmt.Errorf("Try request with a positive number")
			}
		} else {
			// System error
			return 0, err
		}
	}

	return res.GetSquareNumber(), nil
}
