package main

import "fmt"

// function defined outside main func
func add(x int, y int) int {
	return x + y
}

func DefiningFunction(x, y int) {

	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Printf("add(%v, %v): %v\n", x, y, add(x, y))
	fmt.Printf("sub(%v, %v): %v\n", x, y, sub(x, y))
}
