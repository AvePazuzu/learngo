package main

import (
	"fmt"
)

func main() {

	var (
		names     = [...]string{"Name1", "Name2", "Name3"}
		distances = [5]int{123, 124, 125, 126, 127}
		data      [5]byte
		ratios    [1]float32
		alives    [4]bool
		zero      [0]byte
	)

	for i := 0; i < len(names); i++ {
		fmt.Printf("%q \n", names[i])
	}

	for _, v := range distances {
		fmt.Printf("%d \n", v)
	}

	// Print values & Types:
	fmt.Printf("names:     %#v \n", names)
	fmt.Printf("distances: %#v \n", distances)
	fmt.Printf("data:      %#v \n", data)
	fmt.Printf("ratios:    %#v \n", ratios)
	fmt.Printf("alives:    %#v \n", alives)
	fmt.Printf("zero:	   %#v \n", zero)

	// // Print only Types:
	// fmt.Printf("names:     %#T \n", names)
	// fmt.Printf("distances: %#T \n", distances)
	// fmt.Printf("data:      %#T \n", data)
	// fmt.Printf("ratios:    %#T \n", ratios)
	// fmt.Printf("alives:    %#T \n", alives)
	// fmt.Printf("zero:	   %#T \n", zero)

	// // Print only the Values:
	// fmt.Printf("names:     %q \n", names)
	// fmt.Printf("distances: %v \n", distances)
	// fmt.Printf("data:      %v \n", data)
	// fmt.Printf("ratios:    %.2f \n", ratios)
	// fmt.Printf("alives:    %t \n", alives)
	// fmt.Printf("zero:	   %d \n", zero)

}
