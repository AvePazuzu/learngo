package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Give some Numbers!")
		return
	}
	
	toTest:
	for _, v := range args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}	
		
		num := make([]int, n)
		for i, v := range num {
			v = i+1
			num[i] = v 
		}	
		
		for _, v := range num[1:len(num)-1] {
			if num[len(num)-1]%v == 0 {			
				continue toTest		
				}
			}				
			fmt.Printf("%d  ", num[len(num)-1])	
		}
		fmt.Println()	
}