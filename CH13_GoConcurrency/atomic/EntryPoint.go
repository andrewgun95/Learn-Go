package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	const n = 100
	var count int64 = 0

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			// read and write done atomically
			atomic.AddInt64(&count, 1)
			fmt.Println("C :", atomic.LoadInt64(&count))

			wg.Done() // done i-th of n
		}()
	}

	wg.Wait() // wait for n
	fmt.Println("Count :", count)
}
