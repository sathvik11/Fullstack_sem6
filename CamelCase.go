package main

import (
	"fmt"
	"unicode"
)

func isCamelCase(s string) bool {

	if len(s) > 0 && unicode.IsUpper(rune(s[0])) {
		return false
	}

	hasUpper := false
	for _, r := range s {
		if unicode.IsLower(r) {
			hasUpper = true
		} else if hasUpper {
			return true
		}
		for _, r := range s {
			if r == '_' || unicode.IsSpace(r) {
				return false
			}
		}
	}
	return hasUpper
}

func main() {
	test1 := "sathvikAndela"
	test2 := "Sathvik Andela"
	if isCamelCase(test1) {
		fmt.Printf("%s is in camelcase format.\n", test1)
	} else {
		fmt.Printf("%s is not in camelcase format.\n", test1)
	}
	if isCamelCase(test2) {
		fmt.Printf("%s is in camelcase format.\n", test2)
	} else {
		fmt.Printf("%s is not in camelcase format.\n", test2)
	}
}
