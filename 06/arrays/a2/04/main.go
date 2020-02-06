package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	
	var (
		nums [5]float64
		sum float64
		total float64
	)

	args := os.Args
	if len(args) > 6 || len(args) <= 2 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)")
	}
	
	for i := 1; i < len(args); i++ {
		n , err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		nums[i-1] = n
		sum += n
		total++		
	}
	
	fmt.Printf("Your nums are: %v \n", nums)
	fmt.Printf("The avarage is: %.2f \n", sum/total)

}