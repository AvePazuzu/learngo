package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please provide a Name.")
		return
	}

	name := args[1]

	moods := [...]string{"Bad", "Sad", "Happy", "Glad"}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(moods))
	fmt.Printf("%s is %s! \n", name, moods[i])

}
