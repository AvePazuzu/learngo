package main

import (
	"fmt"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	//  1. Separate the data by the newline ("\n") -> rows
	dat := strings.Split(data, "\n")

	//  2. Separate each row of the data by the separator (",") -> columns

	var dat1 [][]string

	for _, v := range dat {
		i := strings.Split(v, separator)
		dat1 = append(dat1, i)
		// fmt.Printf("%#v \n", i)
	}

	// fmt.Printf("%#v \n", dat1)

	n := "Location"

	// fmt.Println((dat[0]))

	for i := 0; i < len(dat1); {
		if n == string(dat1[0][0]) {
			fmt.Println(dat1[i][0])
		}
		i++
	}

	// fmt.Println(dat)l
}
