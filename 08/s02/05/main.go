package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	head := strings.Split(header, separator)

	// fmt.Println(head)
	fmt.Println()

	for i := range head {
		fmt.Printf("%-15s", head[i])
	}
	fmt.Println()

	fmt.Println(strings.Repeat("=", 75))
	// fmt.Println(data)

	dat := strings.Split(data, "\n")

	var (
		locs  []string
		size  []int
		beds  []int
		baths []int
		price []int
	)

	for i := range dat {
		d := strings.Split(dat[i], separator)
		locs = append(locs, d[0])
	}

	for i := range dat {
		d := strings.Split(dat[i], separator)
		s, _ := strconv.Atoi(d[1])
		size = append(size, s)
	}

	for i := range dat {
		d := strings.Split(dat[i], separator)
		b, _ := strconv.Atoi(d[2])
		beds = append(beds, b)
	}

	for i := range dat {
		d := strings.Split(dat[i], separator)
		b, _ := strconv.Atoi(d[3])
		baths = append(baths, b)
	}

	for i := range dat {
		d := strings.Split(dat[i], separator)
		p, _ := strconv.Atoi(d[4])
		price = append(price, p)
	}

	fmt.Printf("\n%-15s%-15d%-15d%-15d%-15d \n", locs[0], size[0], beds[0], baths[0], price[0])
	fmt.Printf("\n%-15s%-15d%-15d%-15d%-15d \n", locs[1], size[1], beds[1], baths[1], price[1])
	fmt.Printf("\n%-15s%-15d%-15d%-15d%-15d \n", locs[2], size[2], beds[2], baths[2], price[2])
	fmt.Printf("\n%-15s%-15d%-15d%-15d%-15d \n", locs[3], size[3], beds[3], baths[3], price[3])

	fmt.Println()
}
