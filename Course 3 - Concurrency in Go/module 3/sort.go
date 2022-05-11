package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func bubbleSort(sli []int64, c chan []int64) {
	for out := 0; out < len(sli); out++ {
		for in := 0; in < len(sli)-out-1; in++ {
			if sli[in] > sli[in+1] {
				swap(sli, in)
			}
		}
	}
	fmt.Println("==============")
	for _, s := range sli {
		fmt.Printf(" %d", s)
	}
	fmt.Println("")
	if c != nil {
		c <- sli
	}
}

func swap(sli []int64, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func main() {

	fmt.Print("Enter a number with space between:  ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := strings.TrimSpace(stdin.Text())

	sli := make([]int64, 0, 0)

	for _, snum := range strings.Split(input, " ") {
		num, err := strconv.ParseInt(snum, 10, 64)
		if err == nil {
			sli = append(sli, num)
		}
	}

	len := len(sli)
	size := int(math.Round(float64(len) / 4))
	var sub1 = make([]int64, 0, 10)
	var sub2 = make([]int64, 0, 10)
	var sub3 = make([]int64, 0, 10)
	var sub4 = make([]int64, 0, 10)
	c := make(chan []int64)
	var j int
	var x = 0
	for i := 0; i < len; i += size {
		j += size

		go bubbleSort(sli[i:j], c)
		if x == 0 {
			sub1 = <-c
		}
		if x == 1 {
			sub2 = <-c
		}
		if x == 2 {
			sub3 = <-c
		}
		if x == 3 {
			sub4 = <-c
		}
		x++
	}
	time.Sleep(1 * time.Second)
	var gen []int64
	gen = append(append(append(sub1, sub2...), sub3...), sub4...)
	bubbleSort(gen, nil)

}
