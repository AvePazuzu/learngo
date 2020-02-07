package main

import (
	"fmt"
)

func main() {
	var (
		names     []string
		distances []float32
		data      []byte
		ratios    []float32
		alives    []bool
	)

	names = []string{}
	distances = []float32{}
	data = []byte{}
	ratios = []float32{}
	alives = []bool{}

	fmt.Printf("names		:%#T 	%d %t \n", names, len(names), names == nil)
	fmt.Printf("distances	:%#T 	%d %t \n", distances, len(distances), distances == nil)
	fmt.Printf("data		:%#T	%d %t \n", data, len(data), data == nil)
	fmt.Printf("ratios		:%#T 	%d %t \n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives		:%#T 	%d %t \n", alives, len(alives), alives == nil)

}
