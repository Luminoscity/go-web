package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Input a series of integers")
	userArray := readNumbers()

	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go partialSort(userArray, 0, 4, &waitGroup)
	go partialSort(userArray, 1, 4, &waitGroup)
	go partialSort(userArray, 2, 4, &waitGroup)
	go partialSort(userArray, 3, 4, &waitGroup)

	waitGroup.Wait()
	fmt.Printf("%v -> ", userArray)
	slices.Sort(userArray)
	fmt.Printf("%v\n", userArray)
}

func readNumbers() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	numberStrings := strings.Fields(input)
	var numbers []int
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func partialSort(arr []int, segment, divider int, waiter *sync.WaitGroup) {
	slice := getSlice(arr, segment, divider)
	fmt.Printf("%d/%d %v -> ", segment+1, divider, slice)
	slices.Sort(slice)
	fmt.Printf("%v\n", slice)
	waiter.Done()
}

func getSlice(arr []int, segment, divider int) []int {
	if segment < 0 || divider <= 0 {
		return nil
	}
	arraySize := len(arr)
	segmentSize := arraySize / divider
	if segmentSize == 0 {
		segmentSize = 1
	}
	start := segment * segmentSize
	if start >= arraySize {
		return nil
	}
	end := start + segmentSize
	if segment == divider-1 { // Last segment may take the rest of the array
		end = arraySize
	}
	if end > arraySize {
		end = arraySize
	}
	return arr[start:end]
}
