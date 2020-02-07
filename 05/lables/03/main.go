/*  This program takes inputs from the console,
 * converts them to int if possible and
 * checks the inputs if they are primes. */

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	t1 := time.Now()

	// Conslose input and err handling
	// returns a slice of strings
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Give some Numbers!")
		return
	}
	// labled statement to return to
toTest:
	// each element is converted to int if possible
	for _, v := range args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		// for every successfull to int conversion a
		// an empty slice of len n is created and filled
		// with the values asscending to n
		num := make([]int, n)
		for i, v := range num {
			v = i + 1
			num[i] = v
		}
		// each element of slice is modulo checked
		for _, v := range num[1 : len(num)-1] {
			if num[len(num)-1]%v == 0 {
				continue toTest
			}
		}
		fmt.Printf("%d  ", num[len(num)-1])
	}
	fmt.Println()
	t2 := time.Now()
	elapsed := t2.Sub(t1)

	fmt.Println(elapsed)

}
