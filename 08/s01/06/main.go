package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	n := strings.Split(namesA, ", ")

	if len(n) == len(namesB) {
		fmt.Println("Length is the same")
	}

	sort.Strings(n)
	sort.Strings(namesB)

	fmt.Println(n)
	fmt.Println(namesB)

	for i := range namesB {
		if n[i] != namesB[i] {
			return
		}
	}
	fmt.Println("Equal")

}
