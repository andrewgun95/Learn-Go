package main

import (
	"fmt"
)

func main() {

	// 1.a What is pointers ?
	// Is a value which *points to the address* of another variable
	// *points to the address* means value can be an address

	// 1.b What is address in RAM ?
	// in RAM look likes,
	// addrs : 199 200 201 202
	// boxes : [5] [3] [2] [1]
	// you know the address, you can read and write its content

	// 1.c Relate to variable ?
	// named memory address, instead of remember all possible numbers of address,
	// ex in 1 mb RAM, there is 1024 for each 1 byte, possible memory location (or address) to stored , but how in GBs ?

	// var	 : a   b   c   d
	// addrs : 199 200 201 202
	// boxes : [5] [3] [2] [1]

	// 2. Pointer is a type
	// ebnf : "*" BaseType
	// 2.a BaseType ? the type that pointer are pointing into
	// 2.b Every value of the address of another variable will have Type Pointer

	// 3. Example :
	a := 200
	fmt.Printf("a address is %v and value is %v and type is %T\n", &a, a, a)
	// (1) its a pointer to variable a - store address a (&a)
	b := &a
	fmt.Printf("b address is %v and value is %v and type is %T\n", &b, b, b)

	// 4. What is deference ? get value of the address,
	// address -> value
	// ex :
	// &a -> a
	// (2) adding the a variable
	*b++
	fmt.Println(a)

	// Chaining pointers
	chainPointers()

	// Multiple pointers
	multiplePointers()

	// 5. When to use pointers ?
	// Pointers allow you to share a value stored in some memory location. Use pointers when
	// 5.a You don’t want to pass around a lot of data
	// 5.b You want to change the data at a location

	// Mutating a value using pointers
	numb := 12
	fmt.Println("numb befor", numb)
	sum(&numb, 2) // 12 + 2 = 14
	min(&numb, 3) // 14 - 3 = 11
	mul(&numb, 2) // 11 * 2 = 22
	fmt.Println("numb after", numb)
}

func chainPointers() {
	fmt.Println("Chaining")
	a := 5
	fmt.Printf("a = %v address is %v\n", a, &a)
	b := &a
	fmt.Printf("b = %v address is %v\n", b, &b)
	c := &b
	fmt.Printf("c = %v address is %v\n", c, &c)
	d := &c
	fmt.Printf("d = %v address is %v\n", d, &d)

	fmt.Printf("d dereference 3 times to get value of a = %d\n", ***d)
}

func multiplePointers() {
	fmt.Println("Multiple")
	a := 5
	fmt.Printf("a = %v address is %v\n", a, &a)
	b := &a
	fmt.Printf("b = %v address is %v\n", b, &b)
	c := &a
	fmt.Printf("c = %v address is %v\n", c, &c)

	*b = 2
	fmt.Println("a =", a) // a = 2
	*c += 10
	fmt.Println("a =", a) // a = 12
}

func sum(numb *int, x int) {
	*numb += x
}

func min(numb *int, x int) {
	*numb -= x
}

func mul(numb *int, x int) {
	*numb *= x
}

// NOTE :
// Everything in Go is Pass by Value.
// Drop any phrases and concepts you may have from other languages.
// Pass by reference, pass by copy - forget those phrases.
// “Pass by Value.”
// That is the only phrase you need to know and remember.
// That is the only phrase you should use.
// Pass by Value.
// Everything in Go is Pass by Value.
// In Go, what you see is what you get (wysiwyg).
// Look at what is happening.
// That is what you get.

// Useful reads :
// https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
