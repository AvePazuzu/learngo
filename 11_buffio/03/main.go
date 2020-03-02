package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	rx := regexp.MustCompile(`[^A-Za-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)
	// here next token is each word. Scan() jups from word to word.
	for in.Scan() {
		total++
		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)
		// in value of the map is increased each time an already exsiting word appears
		words[word]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}
