package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for _, num := range nums {
			input <- num
		}
		close(input)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	go func() {
		for res := range output {
			fmt.Println(res)
		}
	}()

	wg.Wait()
}
