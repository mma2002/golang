package main

import (
	"fmt"
	"math"
	"strconv"
)


func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	return func(t float64) float64{
		return a/2 * math.Pow(t, 2) + (v * t) + s
	}
}

func main() {
	var a, v, s, t string
	fmt.Printf("Enter a acceleration " + a)
	fmt.Scan(&a)
	fmt.Printf("Enter a velocity " + v)
	fmt.Scan(&v)
	fmt.Printf("Enter a displacement " + s)
	fmt.Scan(&s)
	fmt.Printf("Enter a time " + t)
	fmt.Scan(&t)

	acceleration, _ := strconv.ParseFloat(a, 64)
	velocity, _ := strconv.ParseFloat(v, 64)
	displacement, _ := strconv.ParseFloat(s, 64)
	time, _ := strconv.ParseFloat(t, 64)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(time))
}
