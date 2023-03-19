package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var str1, str2 string
	fmt.Print("Enter two strings separated by a space: ")
	fmt.Scan(&str1, &str2)
	str1, str2 = swap(str1, str2)
	fmt.Println(str1, str2)
}
