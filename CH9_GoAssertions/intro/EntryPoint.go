package main

import (
	"fmt"
)

type Someone interface {
	doSomeone() string
}

type Something interface {
	doSomething() string
}

type Study struct { // Study implements both Someone and Something
	about      string
	person     string
	difficulty int
}

func (s Study) doSomething() string {
	return fmt.Sprint("Do : ", s.about)
}

func (s Study) doSomeone() string {
	return fmt.Sprint("Do : ", s.person)
}

func main() {
	// Type Assertions

	// Try to access underlying value of the type from an interface value

	// 1. interface value ?
	// value of interface type - dynamic type

	// 2. Syntax :
	// x.(T)
	// x is interface value
	// T is a type

	// 2.a T is not an interface type - dynamic type of x is identical to the type T
	// * T must implement the interface type of x, otherwise is invalid

	// 2.b T is an interface type 	  - dynamic type of x implements the interface T
	// * The underlying type of x also implement the interface T, otherwise is invalid

	// 3. Assertion holds the value stored in x and type T
	// If not, run-time panic occurs

	// 4. Case 1
	// For Ex :
	var s Something = Study{
		about:      "Go Programming",
		person:     "Self",
		difficulty: 3,
	}

	// fmt.Println(s.about)
	// Illegal : can't access field Study from value type interface of Something

	result1 := s.(Study)
	fmt.Println(result1.doSomething(), result1.about)

	// result2 := s.(int)
	// Illegal : int doesn't implement type interface of Something

	// 5. Avoid run-time panic, using comma ok idiom
	// v, ok := x.(T)

	result3, ok3 := s.(Someone)
	if ok3 {
		// Comment line 25 - 27
		fmt.Println(result3, ok3) // result3, is zero-value or nill
	} else {
		fmt.Println(result3.doSomeone(), ok3)
	}
}
