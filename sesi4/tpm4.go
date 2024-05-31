package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		wg.Add(1)
		go IIFE(number)
	}


	wg.Wait()
}


func IIFE(number int) {
	fmt.Println(number)
	wg.Done()
}