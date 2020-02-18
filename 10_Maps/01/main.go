package main

import (
	"fmt"
)

func main() {
	// Hint: Store phone numbers as text

	var phones map[string]string
	// Key        : Last name
	// Element    : Phone number

	var products map[int]bool
	// Key        : Product ID
	// Element    : Available / Unavailable

	var mPhones map[string][]string
	// Key        : Last name
	// Element    : Phone numbers

	var basket map[int]map[int]int
	// Key        : Customer ID
	// Element Key:
	// Key: Product ID Element: Quantity

	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", mPhones)
	fmt.Printf("basket     : %#v\n", basket)
}
