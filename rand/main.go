package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i != guess; {

		i = rand.Intn(guess + 1)
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	//for {
	//	fmt.Printf("%v \n", time.Now().UnixNano())
	//	time.Sleep(2500 * time.Millisecond)
	//}
}
