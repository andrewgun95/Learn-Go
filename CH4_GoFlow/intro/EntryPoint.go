package main

import (
	"fmt"
	"strings"
)

func main() {

	// SECTION 3 : FOR, IF, SWITCH

	// 1. FOR
	// repeated execution of a block; single condition, for "clause", range "clause"

	// 1. single condition : repeated when condition is evaluated true
	// for condition {}
	// for {} // no condition, condition is evaluated true
	// similar to
	// while (true) {}

	// 2. for "clause" : control by condition, has "init" and "post" statement
	// for init;condition;post {}
	// "init" executed once before evaluated condition for first iteration
	// "post" executed after each execution block

	var total = 0
	for i := 0; ; i++ { // no condition, condition is evaluated true
		if i < 10 {
			fmt.Printf("%d\t%#b\t%#x\n", i, i, i)
			total += i
		} else {
			break
		}
	}
	fmt.Printf("Total : %d\n", total)

	j := 0
	for j < 10 { // similar to for j < 10 {}
		j++
	}

	// 2. IF
	// conditional execution of two branches according to the value of a boolean expression.

	// initial statement : a statement before entering a control block
	// for 	  init;condition {}
	// if     init;condition {}
	// switch init;condition {}

	if x := Total(5, 2, 3, 4); x > 10 {
		fmt.Printf("Total %d is greater than 10\n", x)
	} else {
		fmt.Printf("Total %d is less than 10\n", x) // x is accessed ini here
	}

	// x can't be accessed in here
	// fmt.Println(x)

	// 3. SWITCH
	// conditional execution of more than on branches according to the comparison between
	// "case" expressions and an expression inside "switch"

	// 3.a "case" expressions not be constants
	// Ex :
	// switch { // similiar as switch true {}
	// 	case a > b :
	// 		// do something
	// 	case d == c :
	// 		// do something
	// }

	// 3.b Switch Type

	str := "Hi, my name Andrew! and I born in 'Bandung', 1995. My age is 24 this years"
	result := ""
loop:
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case "."[0]:
			break loop // break the "Loop" (with label tag loop)
		default:
			var strChar string = fmt.Sprintf("%c", str[i])
			switch strChar {
			case ",", "!", "'": // multiple case
				break // break the switch statement
			default:
				switch codeChar := int(str[i]); {
				case i == 0:
					result += strChar
				case codeChar >= 65 && codeChar <= 90: // boolean operations
					result += strings.ToLower(strChar)
				default:
					result += strChar
				}
			}
		}
	}
	fmt.Printf("%s \nFirst sentence : %s\n", str, result)

}

func Total(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}
