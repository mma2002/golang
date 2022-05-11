package main

import (
	"fmt"
	"time"
)

/*
	A race condition is when two or more goroutines have access to the same resource,
	such as a variable (count) and attempt to read and write to that resource.

	This program executes two goroutines and call function
	if one thread tries to increase an integer and another thread tries to read it
	Following code demonstrates a simple way to create a race condition
*/
var count int

func race(s string) {
	for i := 0; i < 2; i++ {
		count++
		fmt.Println(s, ":", count)
	}
}

func main() {
	go race("hello")
	go race("hi")
	time.Sleep(1 * time.Second)
}