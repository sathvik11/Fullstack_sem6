package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num1 := rand.Intn(10) + 1
	num2 := rand.Intn(10) + 1

	// Add the two variables together and store the result in a third variable
	sum := num1 + num2

	// Print out the values of the variables and the sum of the two numbers
	fmt.Printf("num1: %d\nnum2: %d\nsum: %d\n", num1, num2, sum)
}
