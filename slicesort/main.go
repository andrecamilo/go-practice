// Go program to illustrate how to check
// the given slice is sorted or not
package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {

	// Creating and initializing
	// a structure
	Author := []struct {
		a_name    string
		a_article int
		a_id      int
	}{
		{"Mina", 304, 1098},
		{"Cina", 634, 102},
		{"Tina", 104, 105},
		{"Rina", 10, 108},
		{"Sina", 234, 103},
		{"Vina", 237, 106},
		{"Rohit", 56, 107},
		{"Mohit", 300, 104},
		{"Riya", 4, 101},
		{"Sohit", 20, 110},
	}

	// Sorting Author by their name
	// Using Slice() function
	sort.Slice(Author, func(p, q int) bool {
		return Author[p].a_name < Author[q].a_name
	})

	// Checking the slice is sorted
	// according to their names
	// Using SliceIsSorted function
	res1 := sort.SliceIsSorted(Author, func(p, q int) bool {
		return Author[p].a_name < Author[q].a_name
	})

	if res1 == true {

		fmt.Println("Slice is sorted by their names")

	} else {

		fmt.Println("Slice is not sorted by their names")
	}

	// Checking the slice is sorted
	// according to their total articles
	// Using SliceIsSorted function
	res2 := sort.SliceIsSorted(Author, func(p, q int) bool {
		return Author[p].a_article < Author[q].a_article
	})

	if res2 == true {

		fmt.Println("Slice is sorted by " +
			"their total number of articles")

	} else {

		fmt.Println("Slice is not sorted by" +
			" their total number of articles")
	}

	// Sorting Author by their ids
	// Using Slice() function
	sort.Slice(Author, func(p, q int) bool {
		return Author[p].a_id < Author[q].a_id
	})

	// Checking the slice is sorted
	// according to their ids
	// Using SliceIsSorted function
	res3 := sort.SliceIsSorted(Author, func(p, q int) bool {
		return Author[p].a_id < Author[q].a_id
	})

	if res3 == true {

		fmt.Println("Slice is sorted by their ids")

	} else {

		fmt.Println("Slice is not sorted by their ids")
	}
}
