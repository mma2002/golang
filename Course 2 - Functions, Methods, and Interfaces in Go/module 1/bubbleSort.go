package main

import (
	"fmt"
	"log"
	"strconv"
)

func bubbleSort(sli []int64) {
	for out := len(sli) - 1; out > 1; out-- {
		for in := 0; in < out; in++ {
			if sli[in] > sli[in+1] {
				swap(sli, in)
			}
		}
	}
}

func swap(sli []int64, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}
func main() {
	var sli = make([]int64, 0, 10)

	for i := 1; i <= 10; i++ {
		var snum string
		fmt.Printf("Enter a number " + strconv.Itoa(i) + ":")
		fmt.Scan(&snum)

		num, err := strconv.ParseInt(snum, 10, 64)

		if err != nil {
			log.Fatalf("Incorrect number")
		}

		sli = append(sli, num)
	}

	bubbleSort(sli)
	for _, s := range sli {
		fmt.Println(s)

	}
}
