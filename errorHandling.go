package main

import "fmt"

func area(x int, y int) (*int, error) { // return a *int to signal no usable value is returned if error
	// return nil if x or y = 0
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area!")
	}

	// if error is not satisfied, valid
	area := x * y
	return &area, nil

}

func CheckError() {
	x := 2
	y := 0

	area, err := area(x, y) // usual to name return value 'err'
	if err != nil {         // by convention, error check comes first in code
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("area(%v,%v) = %v;\n", x, y, *area)

}
