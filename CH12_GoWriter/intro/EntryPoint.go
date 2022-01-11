package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Document struct {
	Title   string
	Content string
	Author  string
}

// type Writer interface {
//		Writer(p []byte) (n int, err error)
// }

// 1.a Specify a method that has signature match in Writer interface
func (d *Document) Write(p []byte) (n int, err error) {
	d.Content += fmt.Sprintf(" %s.", p)
	return len(p), nil
}

func main() {
	// 1. Custom Implementation of Writer Interface

	// 1.a Line 20 - 24
	// 1.b Document is also type Writer interface
	var w io.Writer = &Document{
		Title:   "Introduction",
		Content: "Go Lang is",
		Author:  "Andreas",
	}
	// 1.c Write in a Document
	n, _ := w.Write([]byte("a Static Programming Language"))
	fmt.Println("Write", n, "bytes")

	n, _ = w.Write([]byte("Go Lang is about a type"))
	fmt.Println("Write", n, "bytes")

	// 2. Use Writer Interface in Standard Library

	// Remember : Assertion !
	// expression.(type)
	if doc, ok := w.(*Document); ok {
		fmt.Println("Writer is type *Document")
		fmt.Printf("%+v\n", doc)

		// *File has a method Write(p []byte) (n int, err error), so it's also type Writer interface
		w, _ = os.Create("intro.json") // returns *File and error

		// json.NewEncoder receive type Writer interface as an arguments, so it's mean encode to the file
		encoder := json.NewEncoder(w)
		// Encode (as json) a document to the file
		encoder.Encode(doc)
	}

	// Exercise :
	// https://play.golang.org/p/9Zxckwj9NHd
}
