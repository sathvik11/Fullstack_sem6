package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)
	x, y := split(num)
	fmt.Printf("The values %d and %d add up to %d.\n", x, y, num)
}
