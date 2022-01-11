package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const n = 100
	count := 0

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			// read
			temp := count
			// yield this go routine - can be interrupted by other go routine
			runtime.Gosched()
			// write
			count = temp + 1

			wg.Done() // done i-th of n
		}()
	}

	wg.Wait()                     // wait for n
	fmt.Println("Count :", count) // count will never have result to n, since there is a race condition
}
