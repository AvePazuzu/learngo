package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	_ = books
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Tell me a book title.")
		return
	}
	fmt.Println("Search results:")
	query := strings.ToLower(args[1])

	var found bool
	for _, v := range books {
		if strings.Contains(strings.ToLower(v), query) {
			fmt.Println("+ " + v)
			found = true
		}
	}
	if !found {

		fmt.Printf("We don't have the book: %s \n", query)
	}

}
