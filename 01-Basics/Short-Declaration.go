package main

import "fmt"

func main() {
	// " := " is also known as short declaration operator
	x := 69
	fmt.Println("x = ", x)
	// While re-assigning value, we don't need short declaration
	x = x + 96
	fmt.Println("x = ", x)
}
