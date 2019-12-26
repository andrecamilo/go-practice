// Go program to illustrate the
// concept of the call by Reference
package main

import "fmt"

// function which swap values
// taking the pointer to integer
func swap(x, y *int) int {

	// taking a temporary variable
	var tmp int

	tmp = *x
	*x = *y
	*y = tmp

	return tmp
}

// Callbyreference2 Main function
func Callbyreference2() {

	var f int = 700
	var s int = 900

	fmt.Printf("Values Before Function Call\n")
	fmt.Printf("f = %d and s = %d\n", f, s)

	// call by Reference
	// by passing the address
	// of the variables
	swap(&f, &s)

	fmt.Printf("\nValues After Function Call\n")
	fmt.Printf("f = %d and s = %d", f, s)
}
