package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

type employee struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func main() {

	// JSON (Javascript Object Notation)
	// a. Use to send and receive data via REST API
	// b. Most developer preferred way (rather than XML)
	// c. Human readable - machine transferable

	// 1. What is marshalling ?
	// Process of transforming the memory representation of an object to a data format suitable for storage or transmission

	// 1.a In Go, convert struct into JSON Object

	// 1.b Sturct must define as Exported Struct Field ?
	// Only fields with a capital first letter are visible to external programs/packages like the JSON Marshaller.
	// field name -> object key (in json)

	// 1.c Marshalling features

	// (i) "json" key in the struct field tag can store STRING FORMAT
	//      STRING FORMAT ? customize field name mapping to object key
	// Ex : Field `json:"fieldName1"`

	// (ii) STRING FORMAT followed by a comma-separated list of options
	//    	"omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value
	// Ex : Field `json:",omitempty"`

	// (iii) STRING FORMAT "-" the field always ommitted
	// Ex : Field `json:"-"`

	// 2. What is unmarshalling ? Inverse of marshalling
	// Unmarshal parses the JSON-encoded data and stores the result in the value that pointed (must pointer)

	// 2.a In Go, convert JSON Object into struct

	// 2.b ONLY FIELDS FOUNDS IN THE destination type(struct) will be decoded.
	// if there is a field in the JSON that isn't in the destination it will be ignored.
	// Ex : https://play.golang.org/p/dDgG9FQUG6l

	// 3. CASE

	// 1. Marshal

	// 1.a Try to JSON marshal from struct

	p1 := person{
		First: "Andrew",
		Last:  "Gun",
		Age:   26,
	}

	bs1, err1 := json.Marshal(p1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(bs1))

	// 1.b Try to JSON marshal from pointer

	bs2, err2 := json.Marshal(&p1) // Marshalling the value that pointed into
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Marshaling pointer will get the same result", string(bs2) == string(bs1))

	// 1.c Try to JSON marshal from map

	var wordCounts = map[int]string{
		1: "abc", // Integer keys are converted to Strings
		2: "xyz",
		3: "jkl",
	}

	bs3, err3 := json.Marshal(wordCounts)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(bs3))

	// 1.d Try to JSON marshal from primitives

	bs4, err4 := json.Marshal(5)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(string(bs4))

	// 1.e Try to JSON marshal from interface and embed types

	// https://play.golang.org/p/vrX2FbpTyR-

	// 2. Unmarshall

	bs5, _ := json.Marshal(p1)

	// 2.a Try to JSON unmarshal into struct

	jsonText := string(bs5)
	fmt.Println(jsonText)

	var emp *employee = &employee{}

	err5 := json.Unmarshal([]byte(jsonText), emp)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Printf("%+v\n", *emp) // Different between +v and v ? v only print the values of the struct, but +v print the members (field and value) of the struct

	// 2.b Try to JSON unmarshal into map

	var personMap map[string]interface{}
	err6 := json.Unmarshal([]byte(jsonText), &personMap)
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println(personMap)

	// 3. Sorting
	// Sorting a primitives slices or user-defined collections

	// Algorithm :
	// 3.a Comparison between two values in element collections - determines Sort in Ascending or Descending
	// 3.b Swap operations
	// Ex : Bubble Sort, Merge Sort, Quick Sort, or etc

	// 3.c Case
	// https://play.golang.org/p/AU5MgFvql28

	// 3.d Exercise
	// https://play.golang.org/p/4ntJmMoDx4k
	// https://play.golang.org/p/ZcGOevPCXWF
}

// Overview : Underlying type and conversions
// https://play.golang.org/p/AQ3OV7y49XR
// https://play.golang.org/p/cbOEccFVDTZ

// string in backtick `` is raw literal string - read exactly as UTF-8 string with no escape characters
