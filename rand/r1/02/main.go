package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	bonMes1 = `This is a bonus message!🍀`
	losMes1 = `losMes2`
	losMes2 = `losMes3`
	losMes  = `You lost! Try again!🍀`
	turns   = 1
)

func main() {

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
		if rand.Intn(luckyNum+1) != luckyNum {
			continue
		}
		if i == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			fmt.Printf("🎉  %s \n", bonMes1)
		}
		return

	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println(losMes)
	case 1:
		fmt.Println(losMes1)
	default:
		fmt.Println(losMes2)
	}

}
