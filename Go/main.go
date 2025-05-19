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

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
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

	// Receivers
	addPartitionBar("Receivers")

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

	// Maps & Slices
	addPartitionBar("Maps & Slices")

	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	var mySlice []string
	numbers := []int{1, 2, 3}

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")

	log.Println(mySlice)
	log.Println(numbers)
	log.Println(numbers[0:2])

	// Decision Structures
	addPartitionBar("Decision Structures")
	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// Loops & Ranging
	addPartitionBar("Loops & Ranging")

	for i := 0; i <= 3; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for _, animal := range animals {
		log.Println(animal)
	}

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"

	for animalType, name := range animalsMap {
		log.Println(animalType, name)
	}

	// Interfaces
	addPartitionBar("Interfaces")

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
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
