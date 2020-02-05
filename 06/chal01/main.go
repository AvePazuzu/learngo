package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Please provide a Name & Mood [pos | neg].")
		return
	}

	name := args[1]
	mood := args[2]

	feels := [...][2]string{
		{"Bad", "Sad"},
		{"Happy", "Glad"},
	}
	// varialbles are always initialized with 0 if not explicitely initialized with literal
	// so case for i = 0 is not needed
	var i int
	switch {
	case mood == "pos":
		i = 1
		// default:
		// 	fmt.Println("give mood [pos | neg")
		// 	return
	}

	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(len(feels[0]))

	fmt.Printf("%s is %s \n", name, feels[i][j])

}
