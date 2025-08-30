package main

import "fmt"

// преобразует слайс строк в множество (без дубликатов)
func ToSet(s []string) []string {
	unique := make(map[string]struct{})
	result := []string{}

	for _, v := range s {
		if _, exists := unique[v]; !exists {
			unique[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Множество:", ToSet(input))
}
