package main

import "fmt"

func main() {

	// calling the function
	n := rpf()

	// displaying the value
	fmt.Println("Value of n is: ", *n)

}

// defining function having integer
// pointer as return type
func rpf() *int {
	// taking a local variable
	// inside the function
	// using short declaration
	// operator
	lv := 100

	// returning the address of lv
	return &lv
}
