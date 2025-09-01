package main

import (
	"fmt"
)

// переворачивает строку, корректно обрабатывая Unicode-символы
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := ReverseString(input)
	fmt.Println("Оригинал:", input)
	fmt.Println("Перевёрнутая:", reversed)

	inputEmoji := "😀🌍🌏"
	reversedEmoji := ReverseString(inputEmoji)
	fmt.Println("Оригинал с эмодзи:", inputEmoji)
	fmt.Println("Перевёрнутая с эмодзи:", reversedEmoji)
}
