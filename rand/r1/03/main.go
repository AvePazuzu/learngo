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
	turns   = 5
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Pick valid numbers!")
		return
	}

	luckyNum, err := strconv.Atoi(args[1])
	luckyNum1, err1 := strconv.Atoi(args[2])
	if err != nil || luckyNum <= 0 || err1 != nil || luckyNum1 <= 0 {
		fmt.Println("Input not valid, try again!")
		return
	}

	var maxNum int

	if luckyNum <= luckyNum1 {
		maxNum = luckyNum1
	} else {
		maxNum = luckyNum
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= turns; i++ {
		k := rand.Intn(maxNum + 1)
		if k != luckyNum && k != luckyNum1 {
			continue
		}
		if i == 1 {
			fmt.Printf("🥇 FIRST TIME WINNER!!! %v \n", k)
		} else {
			fmt.Printf("🎉  %s %v \n", bonMes1, k)
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
