package main

import (
	"os"
)

func main() {

	// 1. Compare against other language
	// 1.a In Go, encourage to explicitly check for the errors
	// "When the error occur and how to handle it"

	// For Ex :
	file, err := os.Open("filePath")
	if err != nil {
		// handle the error - open file
	}
	var res []byte
	_, err = file.Read(res)
	if err != nil {
		// handle the error - read file
	}

	// 1.b In Java, implicitly check for the errors, *but label too many errors*
	// "How to handle the error"

	// For Ex :
	// try {
	//
	// 	File file = new File("filePath")
	// 	FileInputStream inputStream = new FileInputStream(file)
	// 	... Reading over characters ...
	//
	// } catch (Exception e) {
	//	if (e instanceof IllegalArgumentException) {
	//		handle the error - open file
	// 	} else if (e instanceof FileNotFoundException) {
	//		handle the error - read file
	// 	}
	// }

	// 2. Why Error Handling, it's difference between other programming language ?
	// 2.a Handling error gracefully - code explain everything about the errors
	// 2.b Avoid labeling too many errors

	// 3. Drawback, too many REPETITIVE CODE
	// but, can be avoided caused - ERRORS are VALUES

	// 3.a Error Handling in Go
	// Go encourage using *multi-value return* to make it easy to report an error

	// Syntax :
	// res, err := func() (res, err) {
	//	// do something
	// } ()
	// if err != nil {
	//  // handle the error
	// }

	// Not in the case, since *error its a value, that can be stored as a variable* - Rob Pike
	// https://blog.golang.org/errors-are-values

	// For Ex : Error Writer
	// In directory : example/

	// 4. Error is a build-in type as an interface

	// type error interface {
	//       Error() string
	// }

	// 4.a Create a custom error type
	// https://play.golang.org/p/IPWf3ZYRk1T

	// 4.b Use a general error, using errors and fmt package (instead of creating one)
	// https://play.golang.org/p/xEg7ABfW94w

	// 5. ALWAYS, ALWAYS, ALWAYS CHECKING FOR THE *ERRORS*
	// Write the code with errors before writing the code without errors
	// Exceptions : for fmt.Println - caused it will caused infinite checking for the errors

	// 6. Error Logging

	// 6.a fmt.Println - standard print

	// 6.b log.Println - print from any target (ex : a file), customize the print format (ex : add line number) - add time and date by default

	// 6.c log.Fataln  - print and perform os.Exit(1)
	// os.Exit
	// 0 - success termination
	// 1 - termination with an error

	// For Ex :
	// Use all print variants
	// https://play.golang.org/p/5HvTQFJtm0E
	// Use file as a print target
	// In directory : /file

	// Fatal will make *defer functions will not being executed*
	// https://play.golang.org/p/Sw-HBBSHhuV
	// https://play.golang.org/p/6p_bHzDIlnZ

	// 6.d log.Panicln - similar to called panic()

	// 7. Panic

	// Panic will stop the execution until reach go-routine which call the execution
	// Panic *defer functions still will executed*

	// For Ex :
	// https://play.golang.org/p/xqRZ9PzyGgA

	// Panic with stacktrace
	// https://play.golang.org/p/U24Tipvswhj
	// Panic with recover
	// https://play.golang.org/p/QNXGKv3ru3G

	// 8. Others Control Flow
	// defer, panic, and recover ones of others (ex : if, for, switch, goto) control flow that provided by Go
	// defer 	is control the execution to execute at the end of surrounding function
	// panic 	is control the execution to terminate the middle of execution and pass around the caller function (stack trace)
	// recover	is control the execution to escape from panic
}

// More Example :
// Revisit - Defer Function
// https://play.golang.org/p/vZejoXP3rV3
// https://play.golang.org/p/V2KldZ3L9Qo
// Defer for Mutex
// https://play.golang.org/p/X8sTHO3AfK5
// Revisit - Named Return Values
// https://play.golang.org/p/bXc312qMFff
// Revisit - Assertion vs Convertion
// Assertion  - assert a given type of interface to the actual type (which is the implemented)
// Convertion - convert from type to underlying type or vice versa
