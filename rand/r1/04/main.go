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

	if len(args) != 2 {
		fmt.Println("Pick valid numbers!")
		return
	}

	luckyNum, err := strconv.Atoi(args[1])
	if err != nil || luckyNum <= 0 {
		fmt.Println("Input not valid, try again!")
		return
	}

	rand.Seed(time.Now().UnixNano())
	var k int 
	for i := 1; i <= turns; i++ {
		if luckyNum <= 10 {
			k = rand.Intn(11)
		} else {
			k = rand.Intn(luckyNum + 1)
		}
		
		fmt.Printf("%v ", k)
		if k != luckyNum {
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
