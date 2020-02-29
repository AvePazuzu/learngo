package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	input := `House        Student Name
---------------------------
gryffindor   weasley
gryffindor   hagrid
gryffindor   dumbledore
gryffindor   lupin
hufflepuf    wenlock
hufflepuf    scamander
hufflepuf    helga
hufflepuf    diggory
ravenclaw    flitwick
ravenclaw    bagnold
ravenclaw    wildsmith
ravenclaw    montmorency
slytherin    horace
slytherin    nigellus
slytherin    higgs
slytherin    scorpius
bobo         wizardry
bobo         unwanted`

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house := args[1]

	split := strings.Split(input, "\n")

	houses := make(map[string][]string, len(split))

	for i := 2; i < len(split); i++ {
		j := strings.Fields(split[i])
		key := j[0]
		val := j[1]
		houses[key] = append(houses[key], val)
	}

	delete(houses, "bobo")

	students := houses[house]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("\n---- %s students ---- \n\n", house)
	for _, v := range clone {
		fmt.Printf("+ %s \n", v)

	}
	fmt.Println()
}
