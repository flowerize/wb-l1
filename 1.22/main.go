package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Пример 1: Работа с int64
	aInt := int64(1<<20) + 1
	bInt := int64(1<<20) + 2

	fmt.Println("=== Использование int64 ===")
	fmt.Printf("a = %d, b = %d\n", aInt, bInt)
	fmt.Println("Сложение:", aInt+bInt)
	fmt.Println("Вычитание:", aInt-bInt)
	fmt.Println("Умножение:", aInt*bInt)
	fmt.Println("Деление:", aInt/bInt)

	// Пример 2: Работа с math/big для очень больших чисел
	aBig := new(big.Int).SetInt64(1 << 60)
	bBig := new(big.Int).SetInt64(1 << 50)

	fmt.Println("\n=== Использование math/big ===")
	fmt.Printf("a = %s, b = %s\n", aBig.String(), bBig.String())

	sum := new(big.Int).Add(aBig, bBig)
	fmt.Println("Сложение:", sum.String())

	diff := new(big.Int).Sub(aBig, bBig)
	fmt.Println("Вычитание:", diff.String())

	prod := new(big.Int).Mul(aBig, bBig)
	fmt.Println("Умножение:", prod.String())

	quotient := new(big.Int).Div(aBig, bBig)
	fmt.Println("Деление:", quotient.String())
}
