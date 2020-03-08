package main

import (
	"bufio"
	"fmt"
	"os"
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

	games := []game{
		{item: item{1, "god of war", 50}, genre: "action adventure"},
		{item: item{2, "x-com 2", 30}, genre: "strategy"},
		{item: item{3, "minecraft", 20}, genre: "sandbox"},
	}

	ids

	fmt.Printf("\nInanc's game store has %d games.\n\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
> list   : lists all the games
> quit   : quits
`)
		if !in.Scan() {
			break
		}

		fmt.Println()

		switch in.Text() {
		case "id":
			fmt.Println("Wrong ID.")
			break

		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		}
	}
}
