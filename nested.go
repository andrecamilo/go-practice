package main

import "fmt"

// Author Creating structure
type Author struct {
	name   string
	branch string
	year   int
}

// HR Creating
// nested structure
type HR struct {
	details Author
}

// NestedInit has a object to be a initialized
func NestedInit() {

	// Initializing the fields of the structure
	result := HR{

		details: Author{"Sona", "ECE", 2013},
	}

	// Display the values
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
