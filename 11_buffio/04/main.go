package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	var pattern string
	if args := os.Args; len(args) == 2 {
		pattern = args[1]
	}

	for in.Scan() {
		// here token is whole line
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}

}
