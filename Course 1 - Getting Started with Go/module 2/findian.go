package main

import ("fmt"
		"bufio"
		"os"
		"strings")

func main() {
	fmt.Printf("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strl := strings.ToLower(scanner.Text())

		if (strings.HasPrefix(strl, "i") && 
			strings.HasSuffix(strl, "n") && 
			strings.Contains(strl, "a")) {
			
			fmt.Println("Found!")
		
		} else {
		
			fmt.Println("Not Found!")
		
		}
		break
	}
}