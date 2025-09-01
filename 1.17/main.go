package main

import "fmt"

// выполняет бинарный поиск в отсортированном слайсе
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	targets := []int{3, 4, 9, 1, 10}

	for _, target := range targets {
		index := BinarySearch(arr, target)
		if index != -1 {
			fmt.Printf("Element %d found by index %d\n", target, index)
		} else {
			fmt.Printf("Element %d not found\n", target)
		}
	}
}
