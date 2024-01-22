package main

import "fmt"

func PrintParity(x int) {
	// if x%2 == 0 {
	// 	fmt.Printf("%v is even.\n", x)
	// } else {
	// 	fmt.Printf("%v is odd.\n", x)
	// }

	// 'if' does not need 'else'
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	fmt.Printf("%v is odd.\n", x)

}

func Shorthand(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	// r is out of scope here; would only persist inside if-else statement
	fmt.Printf("%v is odd.\n", x)
}
