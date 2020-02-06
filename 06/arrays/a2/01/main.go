package main

import (
	"fmt"
	"strings"
)

func main() {

	header := [3]string{"First name:", "Last name:", "Nickname:"}
		
	names := [...][3]string{
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}	
	for i := range header {
		fmt.Printf("%-20s ", header[i])
	}	
	fmt.Println()
	fmt.Println(strings.Repeat("=", 50))
	
	for i := range names {
		n := names[i]
		fmt.Printf("%-20s %-20s %-20s\n", n[0], n[1], n[2])
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()
	

	// Solution suggested by Author:

	// names := [...][3]string{
	// 	{"First Name", "Last Name", "Nickname"},
	// 	{"Albert", "Einstein", "emc2"},
	// 	{"Isaac", "Newton", "apple"},
	// 	{"Stephen", "Hawking", "blackhole"},
	// 	{"Marie", "Curie", "radium"},
	// 	{"Charles", "Darwin", "fittest"},
	// }

	// for i := range names {
	// 	n := names[i]
	// 	fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

	// 	if i == 0 {
	// 		fmt.Println(strings.Repeat("=", 50))
	// 	}

	// }
}