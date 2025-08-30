package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("abc", size)
}

func safeSubstring(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) > maxRunes {
		runes = runes[:maxRunes]
	}
	return string(runes)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = safeSubstring(v, 100)
	fmt.Printf("Длина justString: %d байт, %d символов\n",
		len(justString), utf8.RuneCountInString(justString))
}

func main() {
	someFunc()
}
