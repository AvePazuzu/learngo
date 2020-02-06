package main

import (
	"fmt"
)

func main() {

	// multDimArr := [...][5]string{

	// }

	header := [...]string{"First name:", "Last name:", "Nickname:"}

	for i := 0; i < len(header); i++ {
		fmt.Printf("%6q", header[i])
	}
	fmt.Println()
	fmt.Printf("%-6s %-6s", header[0], header[1])
	fmt.Println()
	fmt.Println("Hallo, Gopher!")

}

//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
