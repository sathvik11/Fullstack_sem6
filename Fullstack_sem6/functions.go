package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var num1, num2 int
	fmt.Print("Enter two numbers separated by a space: ")
	fmt.Scan(&num1, &num2)
	sum := add(num1, num2)
	fmt.Printf("The sum of %d and %d is %d.\n", num1, num2, sum)
}
