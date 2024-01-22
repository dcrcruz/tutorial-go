package main

import "fmt"

var ( // var () so that no need to type var on each
	v0 int
	v1 float64
	v2 bool
	v3 string
)

func negate(x *int) {
	neg := -*x
	*x = neg
}

func main() {

	x := 3
	fmt.Println(x) // prints 3
	negate(&x)
	fmt.Println(x) // prints -3

	//declare pointer
	// var ptr *string

	var Greets = "Hey"            // capitalize to make it public
	greeting := "Hello\nLinkedIn" // preferred way of initializing variables
	f := 34.25                    // float64 by default
	raw := `Hello Goodbye`        // can span multiple lines; ` for raw meaning as is

	fmt.Println(greeting)
	fmt.Println(Greets)
	fmt.Println(f)
	fmt.Println(raw)

	// assign greeting to address pointer
	// cannot go before 'greeting' = undefined
	// ptr = &greeting
	var ptr *string = &greeting // simplified version

	fmt.Printf("0 var type = %T; value= %v;\n", v0, v0) // %T = type of var and %V = value of var
	fmt.Printf("1 var type = %T; value= %v;\n", v1, v1)
	fmt.Printf("2 var type = %T; value= %v;\n", v2, v2)
	fmt.Printf("3 var type = %T; value= %v;\n", v3, v3)

	fmt.Println("address of greeting:", ptr) // value is memory address
	fmt.Println("value:", *ptr)              // value is pointer value

	DefiningFunction(1, 1)

	myString := "Hi"
	fmt.Println(*&myString) // prints "Hi"

	PrintParity(2)

	CheckError()

	Defer()

}
