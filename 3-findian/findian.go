/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("Enter some text: ")
	in := bufio.NewReader(os.Stdin)
	text, err := in.ReadString('\n') //read entire line

	reg := regexp.MustCompile("(?i)^i.*a.*n$") //case-insensitive, starts with I, contains A, ends with N
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	} else {
		text = strings.TrimRight(text, "\r\n") //Trim trailing whitespace
		if reg.MatchString(text) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
