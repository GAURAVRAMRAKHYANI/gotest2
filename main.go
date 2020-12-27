package main

import (
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

// Main function
func main() {

	// Reversing the string.
	str := "Geeks"

	// returns the reversed string.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)

	fmt.Println("!... Hello World ...!")
}
