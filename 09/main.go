// Empty file finder

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Get dir from console
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide a dir!")
		return
	}

	// Store the file objects in files varialbe
	files, err := ioutil.ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Program optimization:
	// To optimize the program the backing array is created only once
	// Therefore the size of the backing array needs to be calculated
	// var total is the capacity of the backing array
	var total int

	// Loop over the files in path return the len of each file
	// and add it to total; +1 as additional place for line separator '\n'
	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}

	// Names var is declared wiht the make func
	// only bytes are written to files
	names := make([]byte, 0, total)

	// Loop over the files, check for size and append the names
	for _, file := range files {
		if file.Size() == 0 {
			n := file.Name()
			names = append(names, n...)
			// '\n' seperator needed to separate the bytes by line
			names = append(names, '\n')
		}

	}

	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
