/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	if i < len(slice)-1 {
		slice[i], slice[i+1] = slice[i+1], slice[i]
	}
}

func BubbleSort(slice []int) {
	size := len(slice)
	for i := range size - 1 {
		swapped := false
		for j := range size - i - 1 {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	fmt.Print("Enter up to 10 integers (with spaces in between): ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n') //read entire line
	if err != nil {
		fmt.Printf("Input error: %v\n", err)
		os.Exit(0)
	}
	input = strings.TrimSpace(input) //Trim trailing whitespace
	numbers := strings.Fields(input)

	var integers []int
	for i := 0; i < 10 && i < len(numbers); i++ {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Printf("Invalid input: %v", numbers[i])
			os.Exit(0)
		}
		integers = append(integers, num)
	}
	if len(numbers) > 10 {
		fmt.Println("Using only the first 10 numbers.")
	}

	BubbleSort(integers)

	for _, num := range integers {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
}
