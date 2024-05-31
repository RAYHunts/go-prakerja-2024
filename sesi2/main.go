package main

import (
	"fmt"
)

func main() {
	var pyramidHeight int = 10

	for i := 1; i <= pyramidHeight; i++ {

		for j := 1; j <= pyramidHeight-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

