package main

// Control Flow
// 1. Sequential
// 2. Looping
// 3. Conditional

import (
	"fmt"
)

// z := "something" // Will not work, explanation : line 58 - 60
var z = "something"

func main() {
	result1, error1 := sum(5, 2, 3, 10)
	// What fmt does ?
	// package.Identifier
	// from that package used Println function
	// an identifier is the name variable, function, constant, or etc
	fmt.Println(result1)
	fmt.Println(error1)
	// the "_" (underscore) is to throw away returns
	n, _ := fmt.Println("The sum of 5, 2, 3, 10 is", result1, "Even number is", (result1%2 == 0))
	fmt.Println(n) // bytes used in print

	// SECTION 1 : VARIABLES

	// Short Declarations

	// Declare a variable with identifier "x" with type int and intialize a value
	x := 40 // Similar : var x int = 40
	// Reassign the value
	x = 70

	var y int = 20
	result2 := x + y // It's an expression with an operator (+) and operands (x, y)

	fmt.Println("Result : ", result2) // 70 + 20

	// 1. All you need to knows

	// 1.a Compiler will know the type from the initial value
	var a = "initial" // type of string

	// 1.b Declare multiple variable using var keywords
	var b, c int = 1, 2

	// 1.c Declare a variable with no explicit initialization, will set zero-valued (value is given a default value)
	// (1) false for booleans, (2) 0 for numeric types, (3) "" for strings, and (4) nil for pointers, functions, interfaces, slices, channels, and maps
	var d int

	fmt.Println("A, B, C, D", a, b, c, d)

	var e float32

	fmt.Println("Default value of float is", e)

	// 1.d Short Declarations only has scope within a block (local scope) (can't declare outside the block)
	// 1.e Declaration using "var" keywords has scope within a package level (global scope) or a block (local scope)
	fmt.Printf("\"z\" variable with value %s is accessable in here\n", foo())
	// NOTE : Best practice: keep scope as “narrow” as possible
	// What is scope ?
	// Where a variable exists and is accessible

	// 1.f Static programming language : a variable is declared to hold a VALUE of CERTAIN TYPE
	// var f = 43
	// f = "Something" // Will not work, explanation : already has a type int

	// 2. Practices
	// Print type of variable
	var aa = "something, something"
	var bb = 12.7
	fmt.Printf("Variable \"f\" has value : %v and type : %T\n", aa, aa)
	fmt.Printf("Variable \"g\" has value : %v and type : %T\n", bb, bb)
	// String literals comparisons
	fmt.Println("Literal 1 : ", `abc` == "abc")
	var cc =`\n
\n`
	fmt.Println("Literal 2 : ", cc == "\\n\n\\n")
	cc = `
`
	fmt.Println("Literal 3 : ", cc == "\n")
	fmt.Println("Literal 4 : ", `"` == "\"")

	// Zero values
	var dd int
	var ee float32
	var ff bool
	var gg string
	var hh []int
	var ii *int
	var jj map[int]string
	fmt.Println("Zero value for integer type :", dd)
	fmt.Println("Zero value for float type :", ee)
	fmt.Println("Zero value for boolean type :", ff)
	fmt.Println("Zero value for string type :", gg)
	fmt.Println("Zero value for slices type :", hh, "same as <nil>", hh == nil)
	fmt.Println("Zero value for pointer type :", ii, "same as <nil>", ii == nil)
	fmt.Println("Zero value for map type :", jj, "same as <nil>", jj == nil)
}

// Variadic Parameters - consumes more than one paramaters 'of any types using empty interface' - every value is type of empty interface
// Ex:
// params ...interface{}
func sum(numbs ...int) (int, string) {
	var size int = len(numbs)
	if size == 0 {
		return 0, "The numbers is empty"
	}

	var result = 0
	for i := 0; i < size; i++ {
		result += numbs[i]
	}
	return result, "<nil>"
}

func foo() string {
	return z // variable "z" can be accessable with any block within this package
}

// NOTE :
// GO - can't have unused variable (otherwise will return error)

// Indentifier - name program entities, such as variables, types, functions, values
// Example :
// Not Valid !
// var 1abc = 12
// Valid !
var abc1 = 12

// Predeclared identifier
// https://golang.org/ref/spec#Predeclared_identifiers
// Examples :
// types - bool, byte, int8, int16, int32, string, etc
// contants - true, value, etc
// zero value - nil
// functions - print, println, append, etc

// Keywords - is reserved (from syntax), can't be used as identifier
// Examples :
// break, case, const, continue, if, func, etc

// String Literals
// Raw string literals (un-interpreted) - character sequences between back quote (`) will not unescaped
// Interpreted string literals - character sequences between double quote (") will unescaped

// Basic data type - primitive data type, some of the data types has built-in support in most languages
// Composite data type - compose with basic data type (or primitive) or even with others composite data type
