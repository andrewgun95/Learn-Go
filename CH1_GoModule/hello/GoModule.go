package hello

import "rsc.io/quote" // Execute go commands, resolves imports by using the specific dependency module versions listed in go.mod
// What if it encounters an import of a package not provided by any module in go.mod ?
// Looks up the module containing that package (this case rsc.io/quote) and adds it to go.mod, using the latest version

func Hello() string {
	return quote.Hello() // return "Hello, world."
}
