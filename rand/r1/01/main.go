package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	bonMes		= `This is a bonus message!🍀`
	losMes		= `You lost! Try again!🍀`
	turns		= 5
)

func main(){

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Pick a valid number!")
		return
	}

	luckyNum, err := strconv.Atoi(args[1])
	if err != nil || luckyNum <= 0 {
		fmt.Println("Not a valid number, try again!")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= turns; i++ {
		if rand.Intn(luckyNum + 1) == luckyNum {
			if i == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Printf("🎉  %s \n", bonMes)
			}
			return
		}	
	}
	fmt.Println(losMes)
}