package main

import (
	"fmt"
	"log"
	"os"
)

func squareRoot(n int) (bool, error) {
	if n < 0 {
		return false, fmt.Errorf("Can't perform square root of %v. It's a negative number", n)
	}

	return true, nil
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Error occurs : %v\n", err)
	}

	// set log output into a file
	log.SetOutput(file)

	_, err = squareRoot(-3)
	if err != nil {
		log.Println(err)
	}
}
