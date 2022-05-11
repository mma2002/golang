package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	food       string
	locomotion string
	noise      string
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}
func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}
func (c Bird) Eat() {
	fmt.Println(c.food)
}
func (c Bird) Move() {
	fmt.Println(c.locomotion)
}
func (c Bird) Speak() {
	fmt.Println(c.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}
func (c Snake) Eat() {
	fmt.Println(c.food)
}
func (c Snake) Move() {
	fmt.Println(c.locomotion)
}

func (c Snake) Speak() {
	fmt.Println(c.noise)
}
var animalMap = make(map[string]Animal)

type myAnimal struct {
	atype Animal
}
var listAnimal = make(map[string]myAnimal)

func init() {
	animalMap["cow"] = Cow{"grass","walk","moo"}
	animalMap["bird"] = Bird{"worms","fly","peep"}
	animalMap["snake"] = Snake{"mice","slither","hsss"}
}

func main() {

	for {
		fmt.Printf("> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			strl := strings.ToLower(scanner.Text())
			input := strings.Fields(strl)

			if input[0] == "newanimal" {
				create(input)
			} else {
				query(input)
			}

			break
		}
	}
}

func create(input []string)  {
	listAnimal[input[1]] = myAnimal{animalMap[input[2]]}
	fmt.Println("Created it!")
}

func query(input []string)  {
	output, ok := listAnimal[input[1]]
	if ok {
		switch input[2] {
		case "eat":
			output.atype.Eat()
		case "move":
			output.atype.Move()
		case "speak":
			output.atype.Speak()

		}
	} else {
		fmt.Println("Not found name: " + input[1])
	}
}


