package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	num := rand.Intn(10) + 1
	fmt.Printf("The square root of %d is %g.\n", num, float64(num))
}
