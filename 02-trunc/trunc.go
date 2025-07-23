/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Enter a floating-point number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
	} else {
		var trunc int = int(number)
		fmt.Printf("%v\n", trunc)
	}
}
