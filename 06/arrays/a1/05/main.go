package main

import (
	"fmt"
)

func main() {
	var (
		names = [...]string{
			"Einstein",
			"Shepard",
			"Tesla",
		}

		books = [5]string{
			"Kafka's Revenge",
			"Stay Golden",
		}
	)
	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
