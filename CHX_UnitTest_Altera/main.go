package main

import (
	"fmt"
	"icp/unittesting/calculator"
)

func main() {
	result := calc.Add(1, 2)
	fmt.Println("1 + 2 =", result)
}