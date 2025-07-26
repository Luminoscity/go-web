/**
 * Author: Tim Ambrose
 * 2025
 */

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

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}

func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	zoo := map[string]Animal{}
	for {
		fmt.Print("> ")
		input := bufio.NewReader(os.Stdin)
		line, err := input.ReadString('\n') //read entire line
		if err != nil {
			continue
		}
		line = strings.TrimSpace(line) //Trim whitespace
		params := strings.Fields(line)
		if len(params) != 3 {
			fmt.Println("Usage: newanimal <name> (cow|bird|snake)\nquery <name> (eat|move|speak)")
			continue
		}
		var command string = strings.ToLower(params[0])
		var name string = params[1]
		var action string = strings.ToLower(params[2])

		switch command {
		case "newanimal":
			switch action {
			case "cow":
				zoo[name] = Cow{}
			case "bird":
				zoo[name] = Bird{}
			case "snake":
				zoo[name] = Snake{}
			default:
				fmt.Println("Valid animals: cow, bird, snake")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			animal, animalExists := zoo[name]
			if animalExists {
				switch action {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Valid actions: eat, move, speak")
				}
			} else {
				fmt.Printf("There is no animal named %v\n", name)
			}
		default:
			fmt.Println("Valid commands: newanimal, query")
		}
	}
}
