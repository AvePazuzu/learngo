package main

import (
	"fmt"
	"os"
	"strings"
)

const words = "Things are getting silly and stuff licke this are"

func main() {

	words2 := strings.Fields(words)

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me Keyword!")
		return
	}
	key := strings.ToLower(args[1])
queries:
	for i, v := range words2 {
		w := strings.ToLower(v)
	search:
		switch key {
		case "and", "or", "the":
			break search

		}

		if w != key {
			continue queries
		} else {
			fmt.Printf("# %d: %s \n", i+1, v)
			continue queries
		}

	}
	// fmt.Println("Nothing found")

	// fmt.Println(words2, key)
}
