package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	// Remember :
	// Can access for a type struct 'Person' within the same package level without importing

	person := &Person{
		Name: "Andrew",
		Age:  21,
		SocialFollowers: &SocialFollowers{
			Youtube: 1000,
			Twitter: 400,
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Marshalling error : ", err)
	}

	// Print out raw protobuf object (as a byte array)
	fmt.Println(data)
	fmt.Printf("%T\n", data)

	// Unmarshal raw protobuf object (as a byte array) into an object
	newPerson := &Person{}

	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Unmarshalling error : ", err)
	}

	// Print a new person
	fmt.Println(newPerson.GetName())
	fmt.Println(newPerson.GetAge())
	fmt.Println(newPerson.SocialFollowers.GetYoutube())
	fmt.Println(newPerson.SocialFollowers.GetTwitter())
}
