package main

import "fmt"

// преобразует слайс в множество (map[int]bool)
func toSet(s []int) map[int]bool {
	set := make(map[int]bool)
	for _, v := range s {
		set[v] = true
	}
	return set
}

// возвращает пересечение двух слайсов
func intersection(a, b []int) []int {
	setA := toSet(a)
	setB := toSet(b)

	var result []int
	for key := range setA {
		if setB[key] {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println("Пересечение:", intersection(A, B))
}
