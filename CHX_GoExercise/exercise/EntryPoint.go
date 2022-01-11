package main

import (
	"fmt"
	"runtime"
)

// package level scope

var something int = 21

func main() {
	// JEDI LEVEL 1

	// #1.a Short declarations
	x := 42
	y := "James Bond"
	z := true
	// #1.b Print
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// Multiple print
	fmt.Printf("My name is %v. I'm %v and will %v years old on the next month\n", y, z, y)
	// Go doesn't support a ternary operation (condition ? exp1 : exp2)
	// https://golang.org/doc/faq#Does_Go_have_a_ternary_form

	// #2. var Keywords
	// Solution : https://play.golang.org/p/XupkMZ1IcdB

	// #3. Print function
	// Solution : https://play.golang.org/p/h7NgDleeZdn

	// #4. Custom Type
	// Solution : https://play.golang.org/p/wRok-lDpTNC

	// #5. Underlying Type
	// Solution : https://play.golang.org/p/6Hqhww1-0aR

	// #6. More practices
	something := 13
	// Print in block level scope, not in package level scope
	fmt.Println("something is", something)

	// Print OS and architecture
	fmt.Printf("Looking into %s ...\nRunning using %s in %s operating system on %s architecture\n",
		runtime.GOROOT(),
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH)

	// JEDI LEVEL 2

	// #1 Print dec, bin, hex
	// Solution : https://play.golang.org/p/1DMrL_UzlsR
	// #2 Operators
	// Solution : https://play.golang.org/p/pNgPZIy5yii
	// #3 Constant
	// Solution : https://play.golang.org/p/0l8KGK62rjp
	// #4 Bit shift
	// Solution : https://play.golang.org/p/3rYNHw0LuI8
	// #5 Using iota
	// Solution : https://play.golang.org/p/Hb3YogCPrgD

	// JEDI LEVEL 3

	// #1 Loop prints
	// Solution : https://play.golang.org/p/bPTbQSO5kt3
	// #2 Nested loop
	// Solution : https://play.golang.org/p/ovRlLQU5jWH
	// #3 Loop condition
	// Solution : https://play.golang.org/p/zUiHbuUCOAo
	// #4 Modulus
	// Solution : https://play.golang.org/p/XXOgdSTULg5

	// JEDI LEVEL 4

	// #1 Underlying array
	// Solution : https://play.golang.org/p/jZ8wvfGbVJ3
	// #2 Slice
	// Solution : https://play.golang.org/p/h1ydmTtg5Vb
	// #3 Slicing a slice
	// Solution : https://play.golang.org/p/XajdHjlJ1o5
	// #4 Appending a slice
	// Solution : https://play.golang.org/p/VcvGjlFNoza
	// #5 Deleting a slice
	// Solution : https://play.golang.org/p/-isYbYTlH2K
	// #6 "make" keyword
	// Solution : https://play.golang.org/p/rkuHTaznR_B
	// #7 Multidimension array
	// Solution : https://play.golang.org/p/RNmf2RGl40C
	// #8 Map and slice
	// Solution : https://play.golang.org/p/X0T7g67V1Eg
	// #9 Adding a map
	// Solution : https://play.golang.org/p/B_kKn4xW5cI
	// #10 Deleting a map
	// Solution : https://play.golang.org/p/ZBisVbRRQYz

	// JEDI LEVEL 5

	// #1 Struct and slice
	// Solution : https://play.golang.org/p/gXA3sI6HAfo
	// #2 Struct, slice, and map
	// Solution : https://play.golang.org/p/QdIVjesnGwU
	// #3 Struct with embed types
	// Solution : https://play.golang.org/p/Zl9cn5cYSbP
	// #4 Anonymous Struct
	// Solution : https://play.golang.org/p/r_yVVfwvINy

	// JEDI LEVEL 6

	// #1 Function Syntax
	// Solution : https://play.golang.org/p/ce2cZQVV7Bf
	// #2 Variadics Parameters, Unfurl
	// Solution : https://play.golang.org/p/VuPant1kGHW
	// #3 Defer keywords
	// Solution : https://play.golang.org/p/LNjfONevA8r, https://play.golang.org/p/T6s7H4-3lI2
	// #4 Methods
	// Solution : https://play.golang.org/p/ilXjbL-CC_7
	// #5 Interface
	// Solution : https://play.golang.org/p/MUJVXzJO60I
	// #6 Anonymous Function
	// Solution : https://play.golang.org/p/qQkqCw5zh38
	// #7 Function as a variable
	// Solution : https://play.golang.org/p/NU7jl5XnFZV
	// #8 Return value as a function
	// Solution : https://play.golang.org/p/QBoR6dbdAdd
	// #9 Callback function
	// Solution : https://play.golang.org/p/Cv6pcAJ_Rav
	// Different between callback and event ?
	// Event    is a function that called when something happen
	// Callback is a function that called by another function which takes from passing into relate function parameters
	// https://stackoverflow.com/a/9652434
	// #10 Closure - enclosure a scope variable in function
	// Solution : https://play.golang.org/p/SvCdnyIVI6i

	// JEDI LEVEL 7

	// #1 Address
	// Solution : https://play.golang.org/p/UqlNC4glUGK
	// #2 Method sets
	// Solution : https://play.golang.org/p/nyqiMaFmUVy

	// JEDI LEVEL 8

	// #1 Json Marshal
	// Solution : https://play.golang.org/p/zYs4Bh24RcF
	// #2 Json Unmarshal
	// Solution : https://play.golang.org/p/VbHeejqPlt5
	// #3 Writer Interface
	// Solution : https://play.golang.org/p/G-_v8MVW8SH
	// #4 Sorting Primitives Types
	// Solution : https://play.golang.org/p/M8FDdKE-t6l
	// #5 Sorting User defined Types
	// Solution : https://play.golang.org/p/bTbkFSpAEX9

	// JEDI LEVEL 9

	// PRACTICE :
	// #1 Perfect String
	// Solution : https://play.golang.org/p/aqDUDqdA4BA
	// #2 Staircase
	// Solution : https://go.dev/play/p/CyrPHdcrcz4
	// Concurrency
	// #3 Go-routines and wait group
	// Solution : https://play.golang.org/p/zQouJ-KT5ga
	// #4 Method set revisit
	// Solution : https://play.golang.org/p/iniUCsTBwZA
	// #5 Go run, build and install revisit
	// sample/test
	// #6 Fizz buzz
	// Solution : https://play.golang.org/p/VriQvs-AZ3J
	// #7 Binary search
	// Solution : https://play.golang.org/p/1eEBMT5nRGL
	// #8 Balance
	// Solution : https://play.golang.org/p/nLfbYYLu1JG
	// #9 Prime Decomposition
	// Solution : https://play.golang.org/p/PqL2-Gl7rhb
	// #10 Character Counts
	// Solution : https://play.golang.org/p/SqEpS2EHFBM
	// #11 Polindrome
	// Solution : https://play.golang.org/p/LKdjcUZscR9
	// #12 Polindrome Numbers
	// Solution : https://play.golang.org/p/_qXcjEPnpo5

	// JEDI LEVEL 10
	// #1 Channel Block
	// Solution : https://play.golang.org/p/dx7bJUWZobO, https://play.golang.org/p/Y84ZI6K3XYB
	// #2 Channel Direction
	// Solution : https://play.golang.org/p/Ye2E1Bu8H16
	// #3 Channel Range - Close
	// Solution : https://play.golang.org/p/BLCrd0AQb3z
	// #4 Channel Select
	// Solution : https://play.golang.org/p/xDU4wGtHSMH
	// #5 Channel Comma OK Idiom
	// Solution : https://play.golang.org/p/pQ-qq8R1C96
	// #6 More Exercise
	// Solution : https://play.golang.org/p/J79TN7G7dDz
	// #7 More Exervice
	// Solution : https://play.golang.org/p/Lhaol81hxE-

	// JEDI LEVEL 11
	// #1 Check errors
	// Solution : https://play.golang.org/p/Y6mWiirCOSD
	// #2 Errors info
	// Solution : https://play.golang.org/p/5JjsAanXHhM
	// #3 Custom errors 1
	// Solution : https://play.golang.org/p/5sSYdo8Q8Xb
	// #4 Custom errors 2
	// Solution : https://play.golang.org/p/0VCXLhuCXe3

	// Jedi Level 12
	// #1 Writing the Documentation
	// Solution : /CH16_Testing/dog

	// Jedi Level 13
	// Pratice
	// #1 Year to Alphabeth
	// Solution : https://go.dev/play/p/wwRW6kbpHXA - Un-finish

}
