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

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter filename: ")
	in := bufio.NewReader(os.Stdin)
	filename, err := in.ReadString('\n') //read entire line
	if err != nil {
		fmt.Printf("Filename error: %v\n", err)
		os.Exit(0)
	}
	filename = strings.TrimSpace(filename) //Trim trailing whitespace

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			names = append(names, Name{fname: parts[0], lname: parts[1]})
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	for _, name := range names {
		fmt.Printf("First: %s \tLast: %s\n", name.fname, name.lname)
	}
}
