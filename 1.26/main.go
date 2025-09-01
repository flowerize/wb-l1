package main

import (
	"fmt"
	"strings"
)

// проверяет, что все символы в строке уникальны (без учёта регистра)
func AllUnique(s string) bool {
	s = strings.ToLower(s)
	charsMap := make(map[rune]bool)
	for _, ch := range s {
		if charsMap[ch] {
			return false
		}
		charsMap[ch] = true
	}
	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"АаБб",
		"привет",
		"😀😀",
		"GoLang",
		"1234567890",
		"!@#$%^&*()",
	}

	for _, test := range tests {
		fmt.Printf("Строка: %q → %t\n", test, AllUnique(test))
	}
}
