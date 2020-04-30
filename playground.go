package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	type computer struct {
		brand string
	}

	var a, b *computer

	fmt.Println(a == b)

	a = &computer{"Apple"}
	b = &computer{"Apple"}
	fmt.Println(a == b)
	fmt.Println(*a == *b, &a == &b)

}

func incrByPtr(a *[3]int) {
	fmt.Printf("%p\n", a)

	for i := range a {
		a[i]++
	}

}

func incr(a [3]int) [3]int {
	var b [3]int
	for i := range a {
		b[i] = a[i] + 1
	}
	return b
}
