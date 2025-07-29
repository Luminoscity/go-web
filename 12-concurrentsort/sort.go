/**
 * Author: Tim Ambrose
 * 2025
 */

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

func Sort(slice *[]int, which int, qty int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%v/%v\tSorting subarray: %v\n", which, qty, *slice)
	slices.Sort(*slice)
}

func MergeSortedSlices(slices [][]int) []int {
	merged := []int{}
	index := make([]int, len(slices)) //where in each slice are we

	// continue until finished merging all elements
	for {
		minSlice := -1 //which slice has the minimal starting element
		//find minimum among start elements in sorted slices
		for sl := range slices {
			if index[sl] < len(slices[sl]) {
				if minSlice == -1 || slices[sl][index[sl]] < slices[minSlice][index[minSlice]] {
					minSlice = sl
				}
			}
		}
		if minSlice == -1 {
			break
		}
		merged = append(merged, slices[minSlice][index[minSlice]])
		index[minSlice]++
	}
	return merged
}

func main() {
	fmt.Print("Enter a series of integers (with spaces in between): ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Input error: %v\n", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(input)
	numbers := strings.Fields(input)

	// convert to integers
	var integers []int
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", numStr)
			os.Exit(1)
		}
		integers = append(integers, num)
	}

	// divide into goroutines
	const divisions int = 4
	n := len(integers)

	var wg sync.WaitGroup
	sortedSlices := make([][]int, divisions)

	baseSize := n / divisions
	remainder := n % divisions

	start := 0
	for div := 0; div < divisions; div++ {
		end := start + baseSize
		if div < remainder {
			end++ // distribute the remainder among the first few slices
		}
		sortedSlices[div] = integers[start:end]
		//begin sorting if there elements to sort
		if len(sortedSlices[div]) > 0 {
			wg.Add(1)
			go Sort(&sortedSlices[div], div+1, divisions, &wg)
		}
		start = end
	}

	wg.Wait()
	finalSorted := MergeSortedSlices(sortedSlices)
	fmt.Println("Sorted array:", finalSorted)
}
