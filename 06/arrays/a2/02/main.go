package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		EUR = iota
		JPY
		GBP
	)

	exRates := [...]float64{
		EUR: 0.88,
		JPY: 0.78,
		GBP: 113.02,
	}

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	amount, err := strconv.ParseFloat(args[1], 10)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}
	for _, v := range exRates {
		fmt.Printf("%f USD are %f \n", amount, v*amount)
	}

	// Suggested soulution:
	// const (
	// 	EUR = iota
	// 	GBP
	// 	JPY
	// )

	// rates := [...]float64{
	// 	EUR: 0.88,
	// 	GBP: 0.78,
	// 	JPY: 113.02,
	// }

	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("Please provide the amount to be converted.")
	// 	return
	// }

	// amount, err := strconv.ParseFloat(args[0], 64)
	// if err != nil {
	// 	fmt.Println("Invalid amount. It should be a number.")
	// 	return
	// }

	// fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
	// fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
	// fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)

}
