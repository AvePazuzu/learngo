package main

import (
	"fmt"
)

func main() {

	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	mPhones := map[string][]string{
		"bowen": []string{"202-555-0179"},
		"dulin": []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": []string{"03489940240", "03489900120"},
	}

	basket := map[int]map[int]int{
		100: map[int]int{617841573: 4, 576872813: 2},
		101: map[int]int{576872813: 5, 657473833: 20},
		102: map[int]int{},
	}

	fmt.Printf("dulin's     : %v\n", phones["dulin"])
	fmt.Printf("879401371   : %v\n", products[879401371])
	fmt.Printf("greco's 2nd : %#v\n", mPhones["greco"][1])
	fmt.Printf("Customer #101 is going to buy %#v from Product ID #576872813.\n", basket[101][576872813])
}
