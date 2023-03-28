package main

import (
	"fmt"
)

// Pass by value function
func passValue(x int) {
	x = 10
}

// Pass by reference function
func passReference(x *int) {
	*x = 10
}
func main() {
	// Pass by value
	value := 5
	passValue(value)
	fmt.Println("Value after pass by value:", value)

	// Pass by reference
	reference := 5
	passReference(&reference)
	fmt.Println("Value after pass by reference:", reference)
}
