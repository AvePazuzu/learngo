// 1. get input from console
// 2. make: string.ToLower()
// 3. Store input as key element in map
// 4. return from program as soon as an input occures twice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	input := make(map[string]int)

	for in.Scan() {
		word := in.Text()
		word = strings.ToLower(word)
		input[word]++

		if input[word] > 1 {
			fmt.Println("TWICE")
			return
		}

	}
}
