package main

import (
	"fmt"
)

func main() {
	s := 2000
	for i := 2000; i <= 5000; i++ {
		fmt.Println("n :", i-s)
		result := yearToAlphabeth(i - s)
		fmt.Printf("Year %v to alphabeth : %s\n", i, result)
	}
}

func yearToAlphabeth(n int) string {
	a := 65
	z := 90
	t := (z - a) + 1 // consider 'Z'
	if n >= t {
		x := 1
		for i := n / t; i > 0; i /= t {
			x *= t
		}
		fmt.Println("x :", x)
		fmt.Println(n / x)
		return fmt.Sprintf("%c", a+(n/x)-1) + yearToAlphabeth(n%x)
	} else {
		return fmt.Sprintf("%c", a+n)
	}
}
