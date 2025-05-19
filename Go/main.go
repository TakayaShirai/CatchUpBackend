package main

import (
	"fmt"
	"log"
)

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
