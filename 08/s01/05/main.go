package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var names []string

	names = []string{}
	names = []string{"Einstein", "Shepard", "Tesla"}

	//-----------------------------------
	var books = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	//-----------------------------------
	// this time, do not change the nums array to a slice
	nums := []int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums)

	//-----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	n := strings.Join(names, " and ")
	fmt.Printf("%q\n", n)

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
