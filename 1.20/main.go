package main

import (
	"fmt"
	"strings"
)

// переворачивает порядок слов в строке
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	i, j := 0, len(words)-1

	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}

func main() {
	input := "god bless women"
	reversed := reverseWords(input)
	fmt.Println("Оригинал:", input)
	fmt.Println("Перевёрнутая:", reversed)
}
