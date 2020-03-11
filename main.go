package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "time"
)

func incr(a int) int {
	// a++
	return a+ 1
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	// t1 := time.Now()
	// t2 := time.Now()
	// elapsed := t2.Sub(t1)
	num := 10
	fmt.Println(incr(num))

}