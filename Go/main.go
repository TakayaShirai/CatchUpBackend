package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string = "Goodbye, cruel world!"
	var i int = 2

	fmt.Println(whatToSay)

	i = 1
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
