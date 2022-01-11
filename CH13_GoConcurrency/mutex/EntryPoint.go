package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mx sync.Mutex

func main() {
	const n = 100
	count := 0

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			mx.Lock()
			// # critical section
			// read
			temp := count
			// yield this go routine - can be interrupted by other go routine
			runtime.Gosched()
			// write
			count = temp + 1
			// # end
			mx.Unlock()

			wg.Done() // done i-th of n
		}()
	}

	wg.Wait()                     // wait for n
	fmt.Println("Count :", count) // count will have result to n
}
