package main

type temp struct {
	name  string
	tempC float64
}

// t temp - receiver: t is the var and temp is type
// tempF - method name
// float64 - return type
func (t temp) tempF() float64 { // method
	return (t.tempC * 9 / 5) + 32

	// same reult as
	// func tempF(c city) float64 {
	// 	return (c.tempC *9 /5) + 32
	// }
	// but this is a function
}
