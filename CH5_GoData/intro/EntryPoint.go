package main

import (
	"fmt"
	"reflect"
)

func main() {

	// SECTION 4 : ARRAY, SLICE, MAP

	// What is data structure ?
	// Is a particular way of organizing data in a computer so that it can be used effectively
	// https://www.geeksforgeeks.org/data-structures/

	// 1. ARRAY - static way (not grow in size)
	// numbered sequence of elements of a SINGLE TYPE
	// numbered ?
	// means has a length (or size) to specify how many elements will stored

	// Syntax :
	// x := [number]type

	// 1.a The length of array is part of the type ?
	var a [5]int
	var b [6]int
	fmt.Printf("Type %T in variable \"a\" is not the same as Type %T in variable \"b\"\n", a, b)
	fmt.Println(fmt.Sprintf("%T", a) == fmt.Sprintf("%T", b))
	// 1.b How to get length of the array ?
	fmt.Printf("Variable \"a\" has %d length of array\n", len(a))
	fmt.Printf("Variable \"b\" has %d length of array\n", len(b))
	// 1.c Array addressed use a zero-base index (from 0 to len(arr) - 1)
	var c [3]int
	c[0] = 21
	// index 1 is skipped, set into a zero-value
	c[2] = 27
	// 1.d Accessing elements within an array
	for i := 0; i < len(c); i++ {
		fmt.Printf("Element with index %d has value %v\n", i, c[i])
	}
	// 1.e Example :
	// one-dimensional array
	var d [32]byte
	fmt.Printf("Type : %T\tLength : %d\n", d, len(d))
	// n := 1000
	// var e [2*n]int - length must a constant expression
	var e [2 * 1000]int
	fmt.Printf("Type : %T\tLength : %d\n", e, len(e))
	// two-dimensional array
	var f [2][2]float64
	fmt.Printf("Type : %T\tLength : %d\n", f, len(f))

	// 1.f All you need to knows
	// - Array is primarily building block for slices (slices has underlying array)
	// - Arrays are values, assign to another array means copy of all the elements
	var g [3]int
	g = c
	fmt.Println("c :", c)
	fmt.Println("g :", g)
	c[1] = 1 // change element in 'c' not effect to 'g'
	fmt.Println("c :", c)
	fmt.Println("g :", g)
	// - Length (or size) of an array is part of the type, ex types [10]int are [20]int are different

	// 2. SLICE - dynamic way (grow in size)
	// sequence of elements of a SINGLE TYPE - can specify using composite literal
	// 2.a composite literal ?
	// construct a bunch of primitive values (or other values) into a new value - each time evaluated
	// composite literal has a key for each value
	// 2.a.(i)   key is interpreted as a field name for struct literals,
	// 2.a.(ii)  an index for array and slice literals, and
	// 2.a.(iii) a key for map literals

	// Syntax :
	// x := type{values}
	// type is underlying type of array [n]type, slice []type (or even struct or map)
	// Example :
	// array with composite literal
	h := [3]int{1, 2, 3}
	fmt.Println("h:", h)
	// slice with composite literal
	i := []int{1, 2, 3}
	fmt.Println("i:", i)

	// 2.b Cools way to accessing elements within an array, slice, or other composite data
	for i, v := range i {
		fmt.Printf("Element with index %d has value %v\n", i, v)
	}

	// 2.c Slicing a slice
	j := []int{23, 6, 20, 80, 100, 99}
	fmt.Println("Slice 1 :", j[:])   // no slice, print underlying array [23, 6, 20, 80, 100, 99]
	fmt.Println("Slice 2 :", j[1:])  // slice from index 1 (inclusive) to the end [6, 20, 80, 100, 99]
	fmt.Println("Slice 3 :", j[1:3]) // slice from index 1 (inclusive) to index 2 (exclusive of 3) [6, 20]
	fmt.Println("Slice 4 :", j[:3])  // slice from the start to index 2 (exclusive of 3) [23, 6, 20]

	for i, v := range j[2:len(j)] { // each slice of slice will create own indexes
		fmt.Printf("Element with index %d has value %v\n", i, v)
	}

	// 2.d Appending a slice
	// build-in function to append (or adding) the elements to the end of the slice and return the result
	// [1, 2, 3] append [4, 5] result [1, 2, 3, 4, 5]

	// Syntax :
	// append(slice []T, element ...T)

	// What is different ?
	// ...T use in function parameters to convert from n arguments to an array
	// T... use in caller when passing an array to convert into n arguments

	k := []int{1, 2, 3}
	fmt.Println("Befor :", k)
	k = append(k, 4, 5, 6)
	fmt.Println("After :", k)
	// the same
	l := []int{7, 8, 9}
	k = append(k, l...)
	fmt.Println("After :", k)

	// 2.e Deleting a slice
	m := []int{90, 91, 92, 93, 94, 95}
	fmt.Println("Befor :", m)
	m = deleteFromTo(m, 5, 5)
	fmt.Println("After :", m)

	// 2.f Allocating memory in slice
	// There are two built-in function to ALLOCATES MEMORY
	// 2.f.(i) new  : doesn't initialize a memory and return the address (or pointer)
	// 2.f.(i) make : does    initialize a memory and return the initialize value (not the address)

	// Example :
	var n *[]int = new([]int)
	fmt.Printf("Address %v\nValue : %v\n", n, *n)
	var o []int = make([]int, 10)
	fmt.Printf("Address %v\nValue : %v\n", &o, o)

	// 2.g Difference between composite literal and make for initializing slice ?
	// both 'composite literal' and 'make' will allocating an array into a memory
	// composite literal : can't specify the capacity
	// make              : can   specify the capacity
	// capacity is important ? out of capacity, array underlying a slice may extend - allocating into a new memory

	arr := make([]string, 3, 6)
	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	fmt.Println("Value :", arr)
	fmt.Printf("Length : %d\tCapacity : %d\n", len(arr), cap(arr)) // size   slice is 3 and capacity is 6
	arr = append(arr, "d")
	fmt.Println("Value :", arr)
	fmt.Printf("Length : %d\tCapacity : %d\n", len(arr), cap(arr)) // resize slice to 4 and capacity to 6
	arr = append(arr, "e", "f", "g")
	fmt.Println("Value :", arr)
	fmt.Printf("Length : %d\tCapacity : %d\n", len(arr), cap(arr)) // resize slice to 7 and capacity to 12

	// 2.h Slices hold references to an underlying array ?
	// Changes the element of slices inside function will be visible in caller
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Befor", slice) //  1, 2, 3, 4, 5
	changeZero(slice)           // change value index 0 to -1
	fmt.Println("After", slice) // -1, 2, 3, 4, 5

	// NOTE :
	// make([]int, 50, 100)
	// the same as
	// new([100]int)[0:50]

	// 3. MAP - fast way
	// unordered group of elements of a SINGLE TYPE
	// unordered ?
	// can be accessed using unique keys (as index)
	// key -> value

	// Syntax :
	// x := map[keyType]valueType{
	//		key : value
	//}
	// key and value seperated by colon

	// 3.a Key must equality support, means not be a function, map, or slice (not dymanic type)

	p := map[string]int{
		"a": 1,
		"b": 0,
		"c": 3,
	}
	fmt.Println("Variable \"p\" is", p)
	fmt.Println("Entry value of key \"a\" in variable \"p\" is", p["a"]) // index as a key
	fmt.Println("Entry value of key \"d\" in variable \"p\" is", p["d"]) // not present, return zero-value of int (value type)

	// 3.b Case if want to distinguish between zero-value and missing entry ?
	// Comma "ok" idiom
	pb, ok1 := p["b"]
	fmt.Printf("Entry value of key \"b\" in variable \"p\" is %d and present %t\n", pb, ok1)
	_, ok2 := p["d"]
	fmt.Printf("Entry value of key \"d\" in variable \"p\" is present %t\n", ok2)

	// 3.c Best practice comma "ok" idiom
	if v, ok := p["c"]; ok {
		fmt.Printf("Found value %d of key \"c\" in variable \"p\"\n", v) // print only found the value
	}

	// 3.d Map hold references to an underlying array ?
	// Changes the entry of map inside function will be visible in caller
	q := map[string]int{
		"x": 12,
		"y": 20,
	}
	fmt.Println("Befor", q)
	changeKey(q, "x") // change value key "x" to -1
	fmt.Println("After", q)

	// 3.e Set implemention using map
	// What is set ?
	// sequence of elements of a SINGLE TYPE, with no duplicate values

	// How to ? implemented as a map with value type bool
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}
	if attended["Jacob"] {
		fmt.Println("Jacob was at the meeting")
	}
	if attended["Ann"] {
		fmt.Println("Ann was at the meeting")
	}

	// 3.f Adding and Deleting element in map
	attended["Jacob"] = true

	// delete(map, "key")
	// No error is thrown if you use a key which does not exist. Use comma ok idiom to verify

	if _, ok := attended["Jacob"]; ok {
		delete(attended, "Jacob")
	}

	for k, v := range attended {
		if v {
			fmt.Printf("%v, ", k)
		}
	}
	fmt.Println("was at the meeting")

	// 4. Struct
	// sequence of named elements of a DIFFERENT TYPE, called field (has a NAME and a TYPE)
	// 4.a Field can be specified (1) explicitly as IdentifierList or (2) implicitly as EmbeddedField

	// IdentifierList, ebnf : Identifier { "," Identifier } Type
	// EmbeddedField , ebnf : [*] Type

	// NOTE :
	// IdentifierList can be assigned (=) a number equals of ExpressionList, ebnf : Expression { "," Expression }
	var pi, pa int = 5, 4
	fmt.Println("pi, pa", pi, pa)

	// 4.b Field name must be UNIQUE
	type contact struct {
		email string
		phone string
	}

	type person1 struct {
		first, last string
		age         int
		contact
	}

	// 4.c EmbeddedField - declare with a type (no explicit field name) - TYPE NAME act as a FIELD NAME
	type person2 struct {
		first string
		age   int
		contact
		// Illegal :
		// contact - field name must be unique in struct
	}

	// 4.d Assign and access a value to/from a variable of underlying type STRUCT
	// 4.d.(i) Using composite literal
	p1 := person1{
		first: "Andrew",
		age:   21,
		contact: contact{
			email: "andreasgunawan95@gmail",
			phone: "089512809831",
		},
	}

	// 4.d.(i) Using selector expression

	// Syntax :
	// x.f
	// where f is field or method identifier (selector) of a type variable "x"
	// f (a) donate a field or method in a type variable "x" or
	//   (b) donate a field or method of a nested embedded field in a type variable "x"

	p1.last = "Gun"
	fmt.Println("My profile is", p1)
	fmt.Println("My name is", p1.first, p1.last, "and my age is", p1.age)

	// 4.e Each field has a depth
	// 4.f Each field has a metadata (tag)
	type node struct {
		depth int // depth = 2
	}
	type branch struct {
		node // depth = 1
	}
	type tree struct {
		branch `0` // depth = 0, tag "0"
	}

	t1 := tree{ // outer type
		branch: branch{ // outer type
			node: node{ // inner type
				depth: 2,
			},
		},
	}

	t2 := tree{
		branch: branch{
			node: node{
				depth: 2,
			},
		},
	}

	fmt.Println("t1 is equal to t2", t1 == t2)
	fmt.Println("t1 is equal to t2", reflect.DeepEqual(t1, t2))
	fmt.Println("Depth field \"depth\" is", t2.branch.node.depth)

	rt := reflect.TypeOf(tree{})
	fb, _ := rt.FieldByName("branch")
	fmt.Println("Depth field \"branch\" is", fb.Tag)

	// 4.f Using selector expression (x.f) will look also the PROMOTED field ?
	// Every field declaration in embedded field (inner type) will got PROMOTED to the outer type
	fmt.Println("Depth of the tree is", t2.depth+1)
	fmt.Println("Depth of the tree is", t2.branch.depth+1) // depth is accessable to the outer types

	// 4.g Anonymous Struct
	employee := struct {
		first, last string
		age         int
	}{}

	employee.first = "Andrew"
	employee.last = "Gun"
	employee.age = 21

	circle := struct {
		x, y   int
		radius float64
	}{
		x:      16,
		y:      20,
		radius: 12.3,
	}

	fmt.Println("Employee :", employee)
	fmt.Println("Circle :", circle)
}

func changeZero(slice []int) {
	slice[0] = -1            // hold reference to underlying array
	slice = append(slice, 6) // pass by value
}

func changeKey(m map[string]int, key string) {
	m[key] = -1 // hold reference to underlying array
}

func deleteFromTo(slice []int, from int, to int) []int {
	if from <= to && from > -1 && to < len(slice) {
		if to == len(slice)-1 { // last index
			return append(slice[:from])
		} else {
			return append(slice[:from], slice[to+1:]...)
		}
	}

	return slice
}

// Useful reads
// https://www.geeksforgeeks.org/structure-equality-in-golang/
