package main

import (
	"fmt"
)

// SECTION 5 : OOP in Go

// 1. GO is all about TYPE

// 1.a SINGLE TYPE
// hold a single value (basic types, int, bool, string, etc) or multiple values (array, slices, map)
// 1.b DIFFERENT TYPES - struct
// hold multiple values

// 1.d Type can be classify to
// 1.d.(i)  PRE-DECLARED TYPE - basic types, int, bool, string, array, slices, map
// 1.d.(ii) USER-DEFINED TYPE - struct (compose the types together), alias (new definition)

// 1.e We can (1) use the TYPE or (2) create the TYPE

// # CASE 1

// (1) use the TYPE (int)
var x, y int = 12, 32 // assign a value

// # CASE 2

// (1) create a TYPE with pre-declared type
type myInt int

// (2) use the TYPE (myInt)
var z myInt

// # CASE 3

// (1) create a TYPE with another TYPE - struct
type person struct {
	first, last string
	age         int
}

// (2) use the TYPE (person)
var a, b person = person{ // assign a value using Composite Literal
	first: "Andrew",
	last:  "Gun",
	age:   21,
}, person{
	// Bad Practice
	"James",
	"Bond",
	32,
}

// # CASE 4

// (1) create a TYPE and use the TYPE
var c = struct {
	x, y int
}{}

// 2. Methods in Go
// Go doesn't have a class, but you can have methods on TYPE
// What is method ?
// It's just a function with a reciever args - between func keyword and method name (identifier)

// Syntax :
// func (r reciever) identifier(paramaters) (returns) { ... }

// 2.a Method in struct

type Circle struct {
	x, y   int
	radius float64
}

// Method 'calcArea' is attached to Circle type
func (c Circle) calcArea() float64 {
	return 3.14 * c.radius * c.radius
}

// Method 'calcArea' is not attached to any type
func calcArea(c Circle) float64 {
	return 3.14 * c.radius * c.radius
}

// 2.b Method in non-struct

type myFloat float64

func (f myFloat) stringTwoDecimal() string {
	return fmt.Sprintf("%.2f", f)
}

// 2.c Can only declare a method with a reciever type in the same package

// 3. Interfaces in Go
// What is an interface ?
// Two things :
// - Method sets
// - It's a type

//////////////////////////////////////////////////////////////////////////

// 3.a METHOD SET

// What is method sets ?
// All declared method have associated with a type

// For Ex :
// A Cat have method sets Eat and MakeNoise

type Cat struct {
	name string
	age  int
}

// Cat is SUPERSET of Animal and LivingThings

func (c Cat) Eat() {
	fmt.Printf("%s %d Eat !\n", c.name, c.age)
}

func (c Cat) MakeNoise() {
	fmt.Printf("%s %d Meow !\n", c.name, c.age)
}

// For Ex :
// Cat is implementation of Animal
type Animal interface {
	MakeNoise() // Animal is SUBSET of Cat
}

// Cat is implementation of LivingThings
type LivingThings interface {
	Eat() // LivingThings is SUBSET of Cat
}

// 3.b An interface can have implementations more than one type
// https://play.golang.org/p/fuCqqxEG-vf

// 3.c Empty Interface

// Interface with no method set
// Remember ?
// - An interface is SUBSET of any method sets in any type
// - If there are no method set in interface, its mean ALL TYPES WILL IMPLEMENTS that interface (EMPTY INTERFACE)

// 3.d A METHOD SET in interface
// It's specify whether a method set in the interface have subset with method sets in the other type
// If have, that type is *implementation of interface*
// *implementation of interface* a variable of interface type can stored any value of that type

// 3.e A METHOD SET in interface can be specify explicitly (using method signature) or embed of other interfaces
// https://play.golang.org/p/8HCVkuTjsPb
// 3.f A METHOD SET in interface is the UNION of the method sets explicitly declare and the method sets in embedded interface

//////////////////////////////////////////////////////////////////////////

// 4. Pointers and Methods
// A pointer type can access the method (or fields) of associated of a value type, or otherwise

// Why it's even possible ?
// 4.a GO runtime by automatically will dereference the caller value (if pointer) to a value (if value is receiver), or otherwise

// 4.b How methods works ?
// 4.b.(i) Receiver the same as Parameters - but the arg was passing is the caller value
// Remember: Everthing in Go is "Passed by Value"

// For Ex :
type Something struct{}

func (s Something) doSomething(command string) { // similar to this : func doSomething(s Something, command string)
}

// 5. Pointers, Methods and Interface
// Interface doesn't explicitly specify if implementation of the type must be pointer or not
// Remember: Method sets of a type determines the interface that type implements

// Interface Value	    Method Receiver
// It's Pointer			Pointer Type	  Legal
// It's Pointer			Value Type 	      Legal
// It's Value			Value Type		  Legal
// It's Value			Pointer Type	  Illegal

// Methods Receivers    Interface Value
// -----------------------------
// (t T)                T and *T
// (t *T)               *T

// For Ex :
// Simplify : https://play.golang.org/p/mXfFHGf5Xkr
// https://play.golang.org/p/F_L_M-IddJ2

// 4.c. Case 1
// For Ex :
// Without pointer : https://play.golang.org/p/k9BjdSH9bXr
// With    pointer : https://play.golang.org/p/JK7fTVDfKW7

func main() {
	// 1. Go is all about TYPE
	fmt.Println(x, y)
	fmt.Println(a, b)
	c.x = 12
	c.y = 24
	fmt.Println(c)

	// 2. Method in GO

	// 2.a Method in struct
	circle := Circle{
		x:      12,
		y:      10,
		radius: 12.5,
	}

	fmt.Printf("circle (%d, %d) has area : %.2f\n", circle.x, circle.y, circle.calcArea())

	// Illegal : is not accessible since that was attached to the TYPE
	// calcArea()

	// Try this instead : is accessible since that not attacched to any TYPE
	calcArea(circle)

	// 2.b Method in non-struct
	numb := myFloat(2.132)
	fmt.Println("My float is", numb.stringTwoDecimal())

	// 3. Interface value is Undetermined

	// 3.a Both variable with type interfaces (Animal and LivingThings) can store value of Cat type
	var animal Animal = Cat{"nate", 2}
	var live LivingThings = Cat{"bob", 1}
	animal.MakeNoise()
	live.Eat()

	// 3.b Variable with type interface, also can store pointer of Cat type
	live = &Cat{"jake", 3}

	// 4. More about GO :
	fmt.Printf("%T\n", z)
	// 4.a z can be converted to underlying type
	uz := int(z)
	fmt.Printf("%T\n", uz)

	// 4.b constant of a kind
	const n = 12 // undeterminate type
	z = n        // n type become myInt
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", n) // n type is int
}

// Q & A

// 1. Is GO an object-oriented language ?
// Go has types and method, an allow OOP styles, but it's NO TYPE HIERARCHY

// In Java,
// class is used interchangeable as type - class can become a type

// 2. Abstraction in GO ?
// Core concept, abstractions in terms of what actions our types can execute

// In Java,
// abstractions in terms of what kind of data our types can hold

// 3. No Implementations in GO ?
// Whether or not a type implement an interface is determine automatically by looking the method signature

// NOTE :
// Named types vs anonymous types

// Anonymous types are indeterminate.
// They have not been declared as a type yet.
// The compiler has flexibility with anonymous types.
// You can assign an anonymous type to a variable declared as a certain type.
// If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion.
// You cannot assign a named type to a different named type.

// Useful reads :
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
// https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver
