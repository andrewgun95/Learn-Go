package main

import (
	"fmt"
	"strings"
)

func main() {
	// SECTION 7 : Functions

	// 1. What is function ?
	// Try to seperate code (as a chunk), can in single file or different package (modules)
	// Purpose : DRY (Don't Repeat Yourself) - Code reusability

	// 2. Function is a type
	// The type donates with the same parameters and result types (the type is equally for the same signature)

	fmt.Printf("%T\n", a)
	typeA := fmt.Sprintf("%T", a)
	typeB := fmt.Sprintf("%T", b)
	fmt.Println("Type func A and B is equals :", typeA == typeB)

	// 3. Syntax :
	// func (r reciever) identifier (parameter(s)) (result(s)) { ... }

	// Note :
	// The same idea with, no result is VOID

	// 4. Everything in Go is pass by VALUE (recieved as a copy - can't change the value inside function)
	str := "Hello, "
	c(str)
	fmt.Println("Outside func :", str)

	// 5. Variadic function
	// Its a function with a variadic parameters - final parameters n of type ...int (or any type)
	variadic(1, 2, 3)

	// 6. All you need to knows
	// 6.a No actual arguments pass to the function, n become nill (see inside the function)
	variadic()
	// 6.b The length and capacity will vary depend on arguments
	variadic(4, 5) // values pass -> new slice created
	xi := []int{7, 8, 9}
	variadic(xi...) // no new slice created
	fmt.Println("Change the old slice :", xi)

	// 7. Defer keyword
	// 7.a A defer statement defers (post-poned) the execution of a function until the surrounding function returns

	// For example :
	helloDefer()

	// 7.b A defer statement of the execution a function will push into the STACK
	//	   and when the surrounding function returns will executed in LIFO orders (pop from the STACK)

	// For example :
	countDefer()

	// 8. Anonymous Function - self executing function
	// 8.a Function without identifier (or name)

	func() { // Declare a function without parameters
		fmt.Println("Hello, Anonymous Function !")
	}() // Call the functin

	func(name string) { // Declare a function with parameters
		fmt.Println("Hello,", name, "!")
	}("Andrew") // Call the function and pass the args

	// Remember : in Go, function is first-class citizen
	// What it's means ?
	// 9. Function type behave like others type
	// (i) Can assign to variable, (ii) Passing to a function, and (iii) Return from a function

	// 9.a Assign function to a variable as a value expression
	f := func(names ...string) string {
		return strings.Join(names, " ")
	}
	fmt.Println("From Function", f("Andrew", "Gun"))

	s := "Andrew Gun"
	fmt.Println("From String", s)

	// 9.b Return function from a function

	// Remember function type : same parameters and same results
	var x func(name string) string = foo() // Call the foo return a func(string) (string)

	fmt.Printf("Typ of \"x\" is %T\n", x)
	sx := x("Andrew")

	fmt.Println("From Return Function", sx)
	fmt.Println("From Return Function", x("Gun"))

	fmt.Println("From Return Function in Function", foo()("Jacob"))

	// 9.c Callback function
	// Passing args as a function
	// https://play.golang.org/p/EJiSgURIxEM

	// 9.d Closure

	// Scoping
	// Local and Global Scope
	// Variable(s) created inside  in function - it's local  to the function
	// * will newly created every time the function is called
	// Variable(s) created outside in function - it's global to the function

	// Lexical scoping
	// 1. Can   access variable(s) inside a function, if variable in outer function - as global variables
	// 2. Can't access variable(s) inside a function from outer function
	// 3. Can't access variable(s) inside a function from other function in the same level of block

	// What is closure ?
	// Stand for "closes over" some local variables

	// Function value that references to local variables from outside its body - in an enclosing function

	// Sample :
	// func(x int) {
	//		return func() int { return x } - enclosing around variable "x"
	// }

	// Thinking "Closure" about this,
	// 1. return func (params) (returns) {} as "freezing the code" (de-active code) - in later use
	// 2. before "freezing the code" must "setup the code" first

	// For Ex :
	// https://play.golang.org/p/MgIt9APMqyE

	// 9.e Recursive
	// Function called itself
	// * Recursive must have "exit condition", to escape from the loop - take care not overflows

	// For Ex:
	// https://play.golang.org/p/Ck3HkHGYOVO
}

func foo() func(name string) string {
	return func(name string) string {
		if len(name) > 0 {
			return name
		}

		return "no name"
	}
}

///////////////////////////////////////////////
// Defer in Go

func helloDefer() {
	defer fmt.Print(", World! ", "Andreas\n") // evaluate args in fmt.Print(args)

	fmt.Print("Hello")
} // fmt.Print(", World!\n") is executed

func countDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, ", ") // prints 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	}

	fmt.Println("done")
} // LIFO - prints 9, 8, 7, 6, 5, 4, 3, 2, 1, 0

///////////////////////////////////////////////

func a(a, b string, c ...int) int {
	return 0
}

func b(a, b string, c ...int) int {
	return 1
}

func c(a string) {
	a = a + "James"
	fmt.Println("Inside func :", a)
}

func variadic(p ...int) {
	t1 := fmt.Sprintf("%T", p)
	t2 := fmt.Sprintf("%T", []int{})
	fmt.Println("The final parameters of function \"variadic\" has type []int", t1 == t2)

	// What is nill ?
	// No underlying array of n slices
	if len(p) == 0 && cap(p) == 0 {
		fmt.Println("p is nill")
	} else {
		fmt.Println("p is have", len(p), "length and", cap(p), "capacity")

		p[0] = -1
	}
}

// List of function signature
// s1.
func s1() {}

// s2.
func s2(x int) int { return 0 }

// s3. Can blank (_) the names
// s3. Can leave parenthesized for one unamed result
func s3(a, _ int, z float32) bool { return false }

// s4.
func s4(a, b int, z float32) bool { return false }

// s5. Can specify ...T in the last parameters - indicate accept zero or more arguments
func s5(prefix string, values ...int) {}

// s6.
func s6(a, b int, z float64, opt ...interface{}) (success bool) { return success }

// s7. Can leave the names (or identifiers)
func s7(int, int, float64) (float64, *[]int) { return 0.0, nil }

// s8. Illegal signature
// func s8(a, a, b string) {} - each names must be unique
