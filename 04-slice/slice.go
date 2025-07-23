/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// empty slice of integers with length 0 and capacity 3
	numbers := make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer (or 'X' to exit): ")
		var input string
		_, err := fmt.Scan(&input)
		// input = strings.TrimSpace(input)

		if strings.EqualFold(input, "X") {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to exit.")
			continue
		}

		numbers = append(numbers, num)
		sort.Ints(numbers)

		fmt.Println("Sorted slice:", numbers)
	}
}
