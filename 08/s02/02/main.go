package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		pizzaTop   []string
		depTimes   []time.Time
		gradYear   []int
		lightState []bool
	)

	pizzaTop = append(pizzaTop, "pepperoni", "onions", "extra cheese")
	depTimes = append(depTimes, time.Now(), time.Now())
	gradYear = append(gradYear, 1998, 2005, 2018)
	lightState = append(lightState, true, false, true)

	fmt.Printf("\nPizza	: %v \n", pizzaTop)
	fmt.Printf("\ntimes	: %v \n", depTimes)
	fmt.Printf("\nYears	: %v \n", gradYear)
	fmt.Printf("\nState	: %v \n", lightState)

}
