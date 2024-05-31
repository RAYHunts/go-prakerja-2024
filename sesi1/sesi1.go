package main

import (
	"fmt"
)

func main() {
	var name string = "John Doe"
	var age int = 20
	var isMarried bool = false
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)
	fmt.Printf("%v\n", isMarried)

	// data type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}