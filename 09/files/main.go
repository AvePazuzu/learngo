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

	// Print file names in dir
	for _, file := range files {
		fmt.Println(file.Mode())

	}

}
