package main

import "fmt"

// function which modifies
// the value
func modify(Z *int) {
	*Z = 70
}

func main() {

	var Z int = 10

	fmt.Printf("Before Function Call, value of Z is = %d", Z)

	// call by Reference
	// by passing the address
	// of the variable Z
	modify(&Z)

	fmt.Printf("\nAfter Function Call, value of Z is = %d", Z)
}
