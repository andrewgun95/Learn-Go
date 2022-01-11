package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	APIPath = "/api/"
)

type Document struct {
	Title   string
	Content string
	Author  string
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Path[len(APIPath):]

	file, _ := os.Open(fileName + ".json")

	// Read JSON from a File
	decoder := json.NewDecoder(file) // Decode - Unmarshal
	var doc *Document = &Document{}
	err := decoder.Decode(doc)
	if err != nil {
		log.Println(err)
	}

	// Write JSON to a Web Response
	encoder := json.NewEncoder(res) // Encode - Marshal
	err = encoder.Encode(doc)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc(APIPath, viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
