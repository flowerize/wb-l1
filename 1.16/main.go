package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Выбираем первый элемент как опорный
	pivot := arr[0]

	left := make([]int, 0)
	right := make([]int, 0)

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 88, 64}
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Отсортированный массив:", quickSort(arr))
}
