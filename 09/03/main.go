package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get dir from console
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a dir!")
		return
	}

	// Store the file objects in files varialbe
	files, err := ioutil.ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var dirs []byte
	for i := range args[1:] {
		// dirs = append(dirs, '\n')
		dirs = append(dirs, args[i+1]...)
		dirs = append(dirs, '\n')
		for _, v := range files {
			if v.IsDir() {

				dirs = append(dirs, '	')
				dirs = append(dirs, v.Name()...)
				dirs = append(dirs, '/')
				// dirs = append(dirs, '\n')
			}
			dirs = append(dirs, '\n')
		}
		dirs = append(dirs, '\n')

	}

	// '0644' is a Unix code for file permissions
	err = ioutil.WriteFile("out.txt", dirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
