package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("My name: AvePazuzu")
	fmt.Println("My Friend's Name: Lamashtu")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}