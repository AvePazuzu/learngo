package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	a := [...]int{1, 2, 3}

	fmt.Println(a)
	incrByPtr(&a)
	fmt.Println(a)

	a = incr(a)
	fmt.Println("a: ", a)

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
