package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n%-25s %s \n", "Domain", "Visits")
	fmt.Printf("%s \n", strings.Repeat("-", 40))

	sum := make(map[string]int)
	var total int
	var doms []string
	for in.Scan() {
		fields := strings.Fields(in.Text())
		dom := fields[0]
		count, _ := strconv.Atoi(fields[1])

		if _, ok := sum[dom]; !ok {
			doms = append(doms, dom)
		}

		sum[dom] += count
		total += count
		// fmt.Printf("%-25s %d \n", vis, vis[dom])

	}
	sort.Strings(doms)
	for _, el := range doms {
		visits := sum[el]
		fmt.Printf("%-25s %d \n", el, visits)

	}
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("%-25s %d \n", "Total", total)
	fmt.Println()

}
