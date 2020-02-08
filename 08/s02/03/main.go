package main

import (
	"fmt"
)

func main() {

	toppings := []string{"black olives", "green peppers"}
	toppings = append(toppings, "onions", "extra cheese")

	var pizza []string

	pizza = append(pizza, toppings...)

	fmt.Printf("pizza       : %s\n", pizza)

}
