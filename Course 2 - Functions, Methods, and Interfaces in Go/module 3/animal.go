package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}
var animalMap = make(map[string]Animal)

func init() {
	animalMap["cow"] = Animal{"grass","walk","moo"}
	animalMap["bird"] = Animal{"worms","fly","peep"}
	animalMap["snake"] = Animal{"mice","slither","hsss"}
}
func main() {

	for {
		fmt.Printf("> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			strl := strings.ToLower(scanner.Text())
			info := strings.Fields(strl)

			v := animalMap[info[0]]
			if info[1] == "eat" {
				v.Eat()
			} else if info[1] == "move" {
				v.Move()
			} else if info[1] == "speak" {
				v.Speak()
			}
			break
		}
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
