package main

import (
	"fmt"
)

func SetBit(num int64, i int, value int) int64 {
	if i < 0 || i > 63 {
		panic("Bit position must be between 0 and 63 for int64")
	}

	mask := int64(1) << i

	switch value {
	case 1:
		return num | mask
	case 0:
		return num &^ mask
	}

	panic("Value must be 0 or 1")
}

func main() {
	var num int64
	fmt.Scanf("%d", &num)
	i := 0
	value := 0

	result := SetBit(num, i, value)
	fmt.Printf("Original: %d (0b%08b)\n", num, num)
	fmt.Printf("Modified: %d (0b%08b)\n", result, result)
}
