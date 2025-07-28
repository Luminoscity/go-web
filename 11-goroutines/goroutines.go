/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"fmt"
	"time"
)

var sharedI int

func f1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("F1 %v\n", sharedI)
		sharedI++
	}
}

func f2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("F2 %v\n", sharedI)
		sharedI++
	}
}

func main() {
	sharedI = 0
	go f1()
	f2()
}
