package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)
	// here next token is each word. Scan() jups from word to word.
	for in.Scan() {
		total++
		// in value of the map is increased each time an already exsiting word appears
		words[in.Text()]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}
