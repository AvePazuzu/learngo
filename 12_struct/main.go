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

	sum := make(map[string]int)
	var total int
	var doms []string
	cnt := 0

	for in.Scan() {
		cnt++
		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("Worng input: [%s] (line #%d) \n", fields[0], cnt)
			return
		}
		dom := fields[0]
		count, err := strconv.Atoi(fields[1])
		if err != nil || count < 1 {
			fmt.Printf("Worng input: %q (line #%d) \n", fields[1], cnt)
			return
		}

		if _, ok := sum[dom]; !ok {
			doms = append(doms, dom)
		}

		sum[dom] += count
		total += count

	}
	fmt.Printf("\n%-25s %s \n", "Domain", "Visits")
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	sort.Strings(doms)
	for _, el := range doms {
		visits := sum[el]
		fmt.Printf("%-25s %d \n", el, visits)

	}

	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("%-25s %d \n", "Total", total)
	fmt.Println(cnt)

}
