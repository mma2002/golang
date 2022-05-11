package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sli = make([]int, 3)

	for {
		var snum string
		fmt.Printf("Enter a number to add into SLICE or X to exit:")
		fmt.Scan(&snum)

		if snum == "X" {
			break
		}
		num, err := strconv.Atoi(snum)

		if err == nil {
			sli = append(sli, num)
		}
		sort.Ints(sli)
		for _, s := range sli {
			if s > 0 {
				fmt.Println(s)
			}

		}
	}

}
