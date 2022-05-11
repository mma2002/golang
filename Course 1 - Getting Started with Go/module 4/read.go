package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Printf("Enter a full path and name of the text file:")
	fmt.Scan(&fileName)
	dat, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open")

	}
	sli := make([]Person, 0)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		name := strings.Fields(scanner.Text())
		sli = append(sli, Person{name[0], name[1]})
	}

	defer dat.Close()
	for _, s := range sli {

		fmt.Println("First name = " + s.fname + " | Last name = " + s.lname)

	}
}
