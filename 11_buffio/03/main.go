package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	//rx := regexp.MustCompile()

	for in.Scan() {
		fmt.Println(in.Text())

	}

}
