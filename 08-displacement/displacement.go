/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"fmt"
	"os"
)

func GenDisplaceFn(acc float64, vel0 float64, dis0 float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*time*time + vel0*time + dis0
	}
}

func main() {
	var acc float64
	fmt.Print("Enter acceleration: ")
	_, err := fmt.Scan(&acc)
	if err != nil {
		fmt.Println("Invalid acceleration value")
		os.Exit(1)
	}

	var vel0 float64
	fmt.Print("Enter intial velocity: ")
	_, err = fmt.Scan(&vel0)
	if err != nil {
		fmt.Println("Invalid velocity value")
		os.Exit(1)
	}

	var dis0 float64
	fmt.Print("Enter intial displacement: ")
	_, err = fmt.Scan(&dis0)
	if err != nil {
		fmt.Println("Invalid displacement value")
		os.Exit(1)
	}

	var time float64
	fmt.Print("Enter time: ")
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Println("Invalid time value")
		os.Exit(1)
	}

	fn := GenDisplaceFn(acc, vel0, dis0)
	fmt.Printf("Final Displacement after %v seconds: %v", time, fn(time))
}
