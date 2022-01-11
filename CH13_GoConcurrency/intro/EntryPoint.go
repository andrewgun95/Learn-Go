package main

import (
	"fmt"
)

// Concurrency vs Parallelism
// Concurrency is dealing a lot of things at once
// Parallesim  is doing   a lot of things at once

// It's about the DESIGN
// 1. Scalable - will work in single concurrent unit (single thread) or more concurrent unit (multi thread)
// 2. Optimize - better performance in more concurrent unit (multi thread)

func main() {
	// Wait group
	// For ex :
	// https://play.golang.org/p/RpPDqJopYiz

	// Race condition
	// For ex :
	// /race/Entrypoint.go

	// Solve - race condition
	// 1. Mutex  - guarantee that only one go-routine accessing a shared variable at a single time
	// For ex :
	// /mutex/Entrypoint.go

	// 2. Atomic - single step and the operation can't be interrupted by other go-routine
	// For ex :
	// /atomic/Entrypoint.go
}

// Solve concurrency using Channels
// https://play.golang.org/p/VzpLpMhkF04
