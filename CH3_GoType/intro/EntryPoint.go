package main

import (
	"fmt"
)

var a int

type myType int

var b myType

func main() {

	// SECTION 2 : TYPES

	a = 14
	fmt.Printf("Variable \"a\" with value %v and type %T\n", a, a)
	b = 20
	fmt.Printf("Variable \"b\" with value %v and type %T\n", b, b)

	// 1. Variable "a" and "b" have a value look like integer
	// Can we reassign "a" with "b" ?
	// a = b
	// Can't, cause variable "a" and "b" have a different types

	// NOTE : Go is about Static Programming Language ?
	// a variable declared to hold a VALUE to certain of TYPE

	// 2. Conversion
	a = int(b) // Convert underlying-type of b (which is an int)
	// Syntax :
	// T(x)
	// T is a type
	// x is an expression that can be converted into T
	fmt.Printf("Variable \"a\" with value %v\n", a)

	// 3. Alternative of Conversion (Alias)
	var c ball = 80
	var d point = 91
	var cType = fmt.Sprintf("%T", c)
	var dType = fmt.Sprintf("%T", d)
	fmt.Printf("Variable \"c\" with value %v has the same type with int : %v\n", c, cType == "int")
	fmt.Printf("Variable \"d\" with value %v has the diff type with int : %v\n", d, dType != "int")

	a = c
	fmt.Printf("Variable \"a\" with value %v\n", a)

	// 4. Rule of Assignation
	type myArrayType []int
	var e myArrayType
	// Variable with unamed type can be assign into a value of new type, if have identical underlying types
	var f []int = e
	fmt.Printf("Variable \"e\" with value %v and type %T\n", f, f)

	// 5. BOOLEAN

	// 6. NUMERICS

	// 7. STRING
	// A string value is a sequence of bytes (possible emptys)
	g := "Hello, Playground"
	bs := []byte(g)
	fmt.Printf("Hello, Playground as %T : %v\n", bs, bs) // by default bytes will print as decimal values
	// byte is alias of uint8 (0-255)

	for i := 0; i < len(g); i++ {
		// # : add leading 0b for binary (%#b), 0 for octal (%#o), 0x or 0X for hex (%#x or %#X)
		fmt.Printf("%#U ", g[i])

		// What is %U ?
		// print in unicode format, U+%04X
		// %#U print in unicode format and the character, U+%04X 'c'
		// where X is upper-case hexadecimal
	}
	fmt.Println("")

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%d %c %d %#x\n", (i + 1), bs[i], bs[i], bs[i])
	}

	// 8. CONSTANTS
	// 8.a Declare once can't be reassign
	const h string = "Hello, Constant"
	const i = false
	fmt.Printf("Constant \"h\" %v %T\n", h, h)
	fmt.Printf("Constant \"i\" %v %T\n", i, i)

	// 8.b Constant declarations
	// Number of identifiers must be equal to the number of expressions
	const pi, size = 3.14, 10

	// Declaration list of constant expression with parenthesized
	// can be omitted from any identifiers assignment, except the first one
	const (
		j, k = 1, 2
		l, m // By default will repeat assignment from the previous one,
		// but the number of identifiers must equal to the previous one
	)
	fmt.Printf("Constant \"j\" %v %T\n", j, j)
	fmt.Printf("Constant \"k\" %v %T\n", k, k)
	fmt.Printf("Constant \"l\" %v %T\n", l, l)
	fmt.Printf("Constant \"m\" %v %T\n", m, m)

	// 8.c iota ? untyped integer constant - it's generator
	// Starting from (0, 1, 2, 3, 4, etc) for each constant declarations
	const (
		Monday = iota // utyped integer constant
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	const (
		Saturday = iota
		Sunday
	)

	fmt.Printf("Constant \"Tuesday\" %v %T\n", Tuesday, Tuesday)
	fmt.Printf("Constant \"Sunday\" %v %T\n", Sunday, Sunday)

	// After using iota, in the first declaration, the iota value will increment by 1 for the next declaration
	const (
		a0 = iota // a0 == 0
		a1 = iota // a1 == 1
		a2 = iota // a2 == 2
	)

	const (
		b0 = 1 << iota // b0 == 1  (iota == 0)
		b1 = 1 << iota // b1 == 2  (iota == 1)
		b2 = 3         // b2 == 3  (iota == 2, unused)
		b3 = 1 << iota // b3 == 8  (iota == 3)
	)

	// 9. Practice
	const (
		bytes  = 1 << (iota * 10)
		kbytes = bytes << (iota * 10) // Shift over 10
		mbytes = bytes << (iota * 10) // Shift over 20
		gbytes = bytes << (iota * 10) // Shift over 30
	)

	fmt.Printf("1  bytes = %T %d \t\t%#b\n", bytes, bytes, bytes)
	fmt.Printf("1 kbytes = %T %d \t\t%#b\n", kbytes, kbytes, kbytes)
	fmt.Printf("1 mbytes = %T %d \t\t%#b\n", mbytes, mbytes, mbytes)
	fmt.Printf("1 gbytes = %T %d \t%#b\n", gbytes, gbytes, gbytes)

	fmt.Printf("%v %T\n", Pi, Pi)
	fmt.Printf("%v %T\n", zero, zero)
	fmt.Printf("%v %T\n", eof, eof)
}

// NOTE :
// 1. Type declarations - bind an identifier to a type
// Alias declarations - from an identifer to a new alias of type
type ball = int // ball and int are identical types
// Type definition - from an identifier to a new type
type point int // point and int are different types

// Useful reads
// https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/types.html

// 2. Type is explicitly declare a type or otherwise (un-type)
// type floating-point 64 constants
const Pi float64 = 3.14159265358979323846

// untyped floating-point constants
const zero = 0.0

// type integer 64 constants
const size int64 = 1024

// untyped integer constants
const eof = -1
