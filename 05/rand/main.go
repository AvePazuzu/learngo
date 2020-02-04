package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := 5
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= guess; i++ {
		
		k := rand.Intn(guess + 1)
		fmt.Println(k)
		if k == 3 {
			fmt.Printf("%d \n", k)
		}		
		continue
	}
	fmt.Println("Test")
	//for {
	//	fmt.Printf("%v \n", time.Now().UnixNano())
	//	time.Sleep(2500 * time.Millisecond)
	//}
}
