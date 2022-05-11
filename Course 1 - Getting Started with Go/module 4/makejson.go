package main

import (
	"encoding/json"
	"fmt"
)
func main() {

	var name, address string
	fmt.Printf("Enter a name:")
	fmt.Scan(&name)
	fmt.Printf("Enter a address:")
	fmt.Scan(&address)

	person := map[string]string{
		"name":    name,
		"address": address,
	}

	barr, _ := json.Marshal(person)
	fmt.Println(string(barr))

}
