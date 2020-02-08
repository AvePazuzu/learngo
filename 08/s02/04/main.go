package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Give numbers!")
		return
	}

	var nums []int

	for i := range args {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			continue
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
