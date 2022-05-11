package main

import ("fmt")


func main() {
	var pointNum float32
	fmt.Printf("Enter a floating point number:")
	fmt.Scan(&pointNum)
	fmt.Println(int(pointNum))

}