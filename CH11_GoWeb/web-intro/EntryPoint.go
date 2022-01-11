package main

import (
	"fmt"
	"log"
	"net/http"
)

// What is HTTP ?
// Hypertext Transfer Protocol
// *Transfer*
// 1. Send Request  to   the Server
// 2. Give Response from the Server

// The process of transfering from request to the Server and response to the Client
// It's via HTTP

// ResponseWriter - writing response to the Client
// * In background process, it will send back to the Client
// Request        - it's data structure that represent Client request
// * It's tell information that gives from Client Request
// What kind of information ?
// Request Parameters, Path Variable, Request Header and Body, etc
//
// For Ex :
// http://example/hello?first="andrew"&last="gun"&age=25

// URL  				: http://example
// Path 				: /hello
// Path Variable        : hello
// Request Parameters   : first="andrew", last="gun", age=25
// Request Body         : content as json, text, xml, etc, by specify content-type to application/json, application/text, application/xml, etc
// Request Header       : user log-in info (ex : token, season), content-type, etc

func handler(res http.ResponseWriter, req *http.Request) {
	// Get path from url, ignoring root path (/)
	urlPath := req.URL.Path[1:]
	fmt.Fprintf(res, "Hi there, I love %s!\n", urlPath)
	fmt.Println(urlPath)
}

func main() {
	// Handle for all request to the root ("/") at index to handler function (above)
	http.HandleFunc("/", handler)

	// Server should listen to port 8080
	err := http.ListenAndServe(":8080", nil) // The function will be BLOCK until the server is terminated

	// ListenAndServe always return an error - when an unexpected error occurs
	if err != nil {
		log.Fatal(err)
	}
}
