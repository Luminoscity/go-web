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

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (an Animal) Eat() {
	fmt.Println(an.food)
}

func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an Animal) Speak() {
	fmt.Println(an.noise)
}

func main() {
	var kingdom map[string]Animal = map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		fmt.Print("> ")
		input := bufio.NewReader(os.Stdin)
		command, err := input.ReadString('\n') //read entire line
		if err != nil {
			continue
		}
		command = strings.TrimSpace(command) //Trim whitespace
		params := strings.Fields(command)
		if len(params) != 2 {
			fmt.Println("Usage: <animal> <action>\n(cow, bird, snake) (eat, move, speak)")
			continue
		}
		var animal, validAnimal = kingdom[strings.ToLower(params[0])]
		var action = strings.ToLower(params[1])

		if validAnimal {
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
			fmt.Println("Valid animals: cow, bird, snake")
		}
	}
}
