//When storage is allocated for a variable, either through a declaration or a call of new,
//or when a new value is created, either through a composite literal or a call of make,
//and no explicit initialization is provided, the variable or value is given a default value.
//Each element of such a variable or value is set to the zero value for its type:
//	false for booleans,
//	0 for numeric types,
//	"" for strings,
//	nil for pointers, functions, interfaces, slices, channels, and maps.

package main

import "fmt"

//y := 59		// Throws Error as short declarations can't be used outside function declarations
var y = 59 // while var can be used outside function declarations
var z int  // Declare a variable with identifier as 'z' of type int having default value = 0

func main() {

	x := 49

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	y = 69
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)
	z = 123
	fmt.Println("z = ", z)

}
