package main

import (
	"fmt"
)

func main() {
	// 1. What is channel ?
	// Channel is a HIGH LEVEL concurrency

	// 1.a GO Proverb :
	// "Do not communicate by sharing memory; instead, share memory by communicating"

	// What it's mean ?
	// "Do not communicate by sharing memory"
	// (i) Communicate don't exist
	// (ii) Try to communicate by sharing memory
	// "Share memory by communicating"
	// (i) Communicate do    exist
	// (ii) Share memory through communication

	// 1.b Concurrent Programming : Go vs Others

	// Case Example :

	// You and your friend Bob live in different cities, and one day, around dinner time, you start to
	// wonder who ate lunch first that day, you or Bob.
	// How to make sure tomorrow it's the day that you will eat first before Bob ?
	// Solution :
	// #1 Approach - Give Bob a simple instruction not to eat until you call him
	// #2 Approach - Give Bob a message that arrive when you finish the ate

	// 1.c Channels allow us to pass values between goroutines
	// Go encourages a different approach in which **shared values** are passed around on channels
	// In fact, never actively shared by separate threads of execution
	// (i)  Only one goroutine has access to the value at any given time.
	// (ii) Data races cannot occur, by design.

	// 2. Syntax :
	// Channels perform two operation
	// *Send* into the channel
	// c <- value
	// *Receive* from the channel
	// <- c
	// For Ex :
	// https://play.golang.org/p/XXZ454rW597 - Recieve Block
	// https://play.golang.org/p/vx5iGJbwYRr - Send    Block

	// 3. Channel is BLOCK ! - Image : Like runners in a relay race have to pass / receive the baton to each other at the same time
	// For Ex :
	// c := make(chan int)

	// c <- 42 // will block main goroutine
	// fmt.Println(<- c)

	// Solution :
	// Add a new goroutine
	// https://play.golang.org/p/95mHvZHxDbv

	// Add a buffer channel
	// https://play.golang.org/p/50Us3wss8pU

	// What is buffer ? It's maximum amount of value that can keep in the channel

	// 4. Bi-directional Channel vs One-directional Channel

	// Bidirectional  Channel
	c := make(chan int) // Send & Receive
	// Onedirectional Channel
	cr := make(<-chan int) // Receive only
	cs := make(chan<- int) // Send    only

	// 4.a How to know recieve-only and send-only channel ? Read from LEFT to RIGHT

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// 4.b Bi-directional Channel can be constrained into One-directional Channel (either recieve-only or send-only)
	//     using Assignment and Explicit Conversion
	//     For Ex : https://play.golang.org/p/WsxGiXADE_Q

	// 4.c Exercise : https://play.golang.org/p/h5JjiOyZzHk

	// 5. Range over the Channel will pull of the value of the Channel until it closed
	// For Ex : https://play.golang.org/p/fNp7UR_pBbV

	// 6. Select can use to receive a value from a Multiple Channel at once

	// Similar with switch but,
	// It's decision is based on ability to communicate rather than equals values - whether channel is ready to receive or not
	// For Ex :

	o, e, q := make(chan int), make(chan int), make(chan int)
	select {
	case <-o:
		// Execute this statement if channel 'o' is ready to receive
		fmt.Println("Channel 'o' is ready")
	case <-e:
		// Execute this statement if channel 'e' is ready to receive
		fmt.Println("Channel 'e' is ready")
	case <-q:
		// Execute this statement if channel 'q' is ready to receive
		fmt.Println("Channel 'q' is ready")
	default:
		// Neither channels is ready
		fmt.Println("Channels not ready")
	}
}

// More Example :
// Revisit - Try Assignment and Conversion
// https://play.golang.org/p/DUywmmzQyEy
// Revisit - Try Fallthrough
// https://play.golang.org/p/9zqJK2CAasb
// Channel Buffered :
// 1. https://play.golang.org/p/j0OtHCeX1Lv
// 2. https://play.golang.org/p/-b5BVCRbKZw
// Channel - Comma OK Idiom
// https://play.golang.org/p/px26NYWiLvr
// Channel - Select Example
// https://play.golang.org/p/YZIzB3sJqSU, https://play.golang.org/p/u_mVZGWcG90
