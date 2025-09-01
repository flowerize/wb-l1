package main

import (
	"fmt"
)

// удаляет элемент по индексу i из слайса
func removeIndex(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, fmt.Errorf("Index out of range")
	}

	copy(s[i:], s[i+1:])

	return s[:len(s)-1], nil
}

func main() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", s)

	newSlice, err := removeIndex(s, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Slice after deletion:", newSlice)
}
