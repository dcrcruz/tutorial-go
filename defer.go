package lessons

import "fmt"

func printMsg(id string, m string) {
	fmt.Printf("[%v]: %v\n", id, m)
}

func Defer() {
	msg := "hi"
	defer printMsg("Defer-0", msg)
	defer printMsg("Defer-1", msg)

	msg = "bye"
	fmt.Println("exit")

	// output
	// 	exit
	// [Defer-1]: hi
	// [Defer-0]: hi

	// defer captured "hi" instead of "bye" since it was the value of msg
	// when defer was used
}
