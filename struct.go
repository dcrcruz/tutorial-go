package main

import "fmt"

type city struct {
	name  string
	tempC float64
}

func NewCity(n string) city {
	return city{
		name: n,
	}
}

func Struct() {
	c := city{
		name:  "London",
		tempC: 7.5, // can be ommitted; no need to provide values for all
	}

	fmt.Println(c.name) // field of struct is accessed using .
}

func PrintCity(p city) {
	fmt.Printf("%v", p)
}
