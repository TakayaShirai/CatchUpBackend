package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// Variables & Functions
	addPartitionBar("Variables & Functions")
	fmt.Println("Hello, World!")

	var whatToSay string = "Goodbye, cruel world!"
	var i int = 2

	fmt.Println(whatToSay)

	i = 1
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)

	// Pointers
	addPartitionBar("Pointers")
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("myString is set to", myString)

	// Types & Structs
	addPartitionBar("Types & Structs")

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName)
	log.Println(user.LastName)
}

func saySomething() (string, string) {
	return "something", "else"
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}

func addPartitionBar(s string) {
	fmt.Println()
	fmt.Println(s, "----------------------------------------------")
}
