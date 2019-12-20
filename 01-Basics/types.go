package main

import "fmt"

var a = 59
var firstName = "Rajat"
var lastName = `Gupta`

func main() {
	fmt.Println("a = ", a)
	fmt.Printf("Type of a = %T\n", a)

	fmt.Println("First Name = ", firstName)
	fmt.Printf("Type of First Name = %T\n", firstName)

	fmt.Println("Last Name = ", lastName)
	fmt.Printf("Type of Last Name = %T\n", lastName)

	//	a = "Hello"		// this will return an error as type re assignment doesn't works in go as it is a dynamic language
}
