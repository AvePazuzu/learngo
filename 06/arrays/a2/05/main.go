package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		nums [5]float64
	)

	args := os.Args

	switch l := len(args); {
	case l < 2:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case l > 6:
		fmt.Println("Sorry. Go arrays are fixed.",
			"So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	// Fill the array wiht the parsed floats:
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			// Skip if error makes n = 0
			continue
		}
		nums[i-1] = n
	}

	fmt.Println(nums)

	/*
		check whether it's the last element or not:
		  i < len(nums)-1
		check whether the next number is greater than the current one, if so, swap it:
		  v > nums[i+1]
	*/

	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				// fmt.Println(nums)
				// fmt.Println(i, v, i+1, nums[i+1])
			}
		}
	}

	fmt.Println(nums)
}
