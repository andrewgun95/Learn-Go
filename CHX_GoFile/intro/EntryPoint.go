package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get absolute path of current directory
	pwd, _ := os.Getwd()

	fmt.Println("Reading a file ...")
	// 1. Read from content in the file into memory as bytes
	data, err1 := ioutil.ReadFile(pwd + "/file.txt")

	if err1 != nil { // If unable to read the file, print the reason
		fmt.Println(err1)
	}

	// If successful to read the file, print the content
	fmt.Printf("Type : %T\nContent : %s\n", data, data)

	// Construct a byte array as content to store within a file
	myData := []byte("Hello, Jerk!")

	fmt.Println("Writing a file ...")
	// 2. Write the content as byte into a file
	err2 := ioutil.WriteFile(pwd+"/myData.data", myData, 0777) // If failed to write the file, return error
	// 0777 is a File Mode contains a bit permissions

	if err2 != nil { // Handle that error
		fmt.Println(err2)
	}

	// If successful to write the file, print the content
	fmt.Printf("Type : %T\nContent : %s\n", myData, myData)
}

// NOTE :
// format-agnostic ?
// read any file (within any format)

// Useful reads
// https://tutorialedge.net/golang/reading-writing-files-in-go/
// https://golang.org/pkg/os/#FileMode
