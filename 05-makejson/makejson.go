/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter name: ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n') //read entire line
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(0)
	}
	name = strings.TrimRight(name, "\r\n") //Trim trailing whitespace

	fmt.Print("Enter address: ")
	address, err := in.ReadString('\n') //read entire line
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(0)
	}
	address = strings.TrimRight(address, "\r\n") //Trim trailing whitespace

	addressBook := map[string]string{}

	addressBook["name"] = name
	addressBook["address"] = address

	json, err := json.Marshal(addressBook)
	fmt.Println(string(json))
}
