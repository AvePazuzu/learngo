package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// var c *computer
	c := &computer{}
	change(c, "apple")
	// fmt.Printf("%p\n", c)
	fmt.Printf("brand: %s\n", c.brand)
}

func change(c *computer, newBrand string) {
	c.brand = newBrand
}
