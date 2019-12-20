package main

import "fmt"

func main() {
	fmt.Printf("Hello World !!") //Return type of Println = number of bytes written, error
	numberOfBytes, err := fmt.Println("Hello World!!")
	fmt.Println(numberOfBytes)
	fmt.Println(err)
	//If we want not to use variable, just tell compiler explicitly by using '_'
	numberOfBytes, _ = fmt.Println("Hello World This is Rajat!!")
	fmt.Println(numberOfBytes)

}
