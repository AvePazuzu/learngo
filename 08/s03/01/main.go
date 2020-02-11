package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	
	//   1. Convert the string to an []int
	splits := strings.Fields(data)	
	var nums []int

	for i := range splits {
		n, _ := strconv.Atoi(splits[i])

		nums = append(nums, n)
	}

	//   2. Print the slice
	fmt.Println(nums)
	
	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	even := nums[:3]
	fmt.Println(even)
	
	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	odd := nums[3:]
	fmt.Println(odd)

	//   5. Slice it for the two numbers at the middle
	mid := nums[2:4]
	fmt.Println(mid)

	//   6. Slice it for the first two numbers
	first := nums[:2]
	fmt.Println(first)

	//   7. Slice it for the last two numbers (use the len function)
	last := nums[len(nums)-2:]
	fmt.Println(last)

	//   8. Slice the evens slice for the last one number
	evenL := even[len(even)-1:]
	fmt.Println(evenL)

	//   9. Slice the odds slice for the last two numbers
	oddL := odd[len(odd)-2:]
	fmt.Println(oddL)




	
	// fmt.Printf("%#v", nums)
	
}