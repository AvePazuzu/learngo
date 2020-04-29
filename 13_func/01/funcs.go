package main

import (
	"fmt"
	"strconv"
)

func newGame(idd, name, price, genre string) {

	ids, err := strconv.Atoi(idd)
	if err != nil {
		fmt.Println("Nix int")
		return
	}

	prices, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("nix int")
		return
	}

	ng := game{item: item{id: ids, name: name, price: prices}, genre: genre}

	fmt.Printf("ID: %d, Name: %q, Genre: %q, Price: %d\n", ng.id, ng.name, ng.genre, ng.price)
	fmt.Println()
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		ng.id, ng.name, "("+ng.genre+")", ng.price)
	fmt.Println()

}
