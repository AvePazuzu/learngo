package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main()  {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Send me some items and I will sort them")
		return
	}
	toSort := args[1:]
	
	sort.Strings(toSort)

	var names []byte

	// Loop over the files, check for size and append the names
	for _, n := range toSort {

		names = append(names, n...)
		// '\n' seperator needed to separate the bytes by line
		names = append(names, '\n')
	
	}

	// '0644' is a Unix code for file permissions
	err := ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	

	
	// fmt.Printf("%#v \n", names)


}