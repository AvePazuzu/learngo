package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {

	// load()
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	// indexByID()
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list    	: lists all the games
  > add 	: adds a new game to list
  > id N   	: queries a game by id
  > quit   	: quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				// printGame()
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}

		case "add":
			if len(cmd) != 5 {
				fmt.Println("Wrong input!")
				continue
			}

			newGame(cmd[1], cmd[2], cmd[3], cmd[4])

			continue

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			// printGame()
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		}
	}
}
