package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var null *computer

	// compare the pointer variable to a nil value, and say it's nil
	if null == nil {
		fmt.Println("null is nil!")
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{"Apple"}

	// put the apple into a new pointer variable
	newApple := apple

	// compare the apples: if they are equal say so and print their addresses
	if newApple == apple {
		fmt.Printf("newApple @: %p\napple @: %p\n", newApple, apple)
	}

	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{"Sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if sony != apple {
		fmt.Printf("sony @: %p\napple @: %p\n", sony, apple)
	}

	// put apple's value into a new ordinary variable
	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple                     : %p %p\n", &apple, apple)
	fmt.Printf("appleVal                  : %p\n", &appleVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == appleVal {
		fmt.Printf("Is equal!\n")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("apple: %v\nappleVal: %v\n", *apple, appleVal)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "birne")
	fmt.Printf("sony new: %v\n", sony.brand)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	ret(sony)

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(c *computer, newBrand string) {
	c.brand = newBrand

}

func ret(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
