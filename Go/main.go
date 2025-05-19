package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/TakayaShirai/CatchUpBackend/Go/helpers"
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

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
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

	// Packages
	addPartitionBar("Packages")

	var newVar helpers.SomeType
	newVar.TypeName = "Some name"
	newVar.TypeNumber = 1
	log.Println(newVar.TypeName, newVar.TypeNumber)

	// Channels
	addPartitionBar("Channels")

	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)

	// JSON
	myJSON := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true	
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	// Unmarshalling: Convert a JSON string into Go structs.
	// Marshalling: Convert Go structs back into a JSON string.
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJSON), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var newSlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false
	newSlice = append(newSlice, m1)

	var m2 Person
	m2.FirstName = "Bob"
	m2.LastName = "West"
	m2.HairColor = "black"
	m2.HasDog = true
	newSlice = append(newSlice, m2)

	newJson, err := json.MarshalIndent(newSlice, "", "    ")

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
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
