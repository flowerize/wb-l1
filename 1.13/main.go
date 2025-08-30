package main

import "fmt"

// через сложение и вычитание
func SwapAddSub(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

// через XOR
func SwapXOR(a, b *int) {
	if a == b { // Избежать обнуления, если a и b указывают на один адрес
		return
	}
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	x, y := 5, 10
	fmt.Printf("До XOR: x=%d, y=%d\n", x, y)
	SwapXOR(&x, &y)
	fmt.Printf("После XOR: x=%d, y=%d\n", x, y)

	a, b := 20, 30
	fmt.Printf("До сложения: a=%d, b=%d\n", a, b)
	SwapAddSub(&a, &b)
	fmt.Printf("После сложения: a=%d, b=%d\n", a, b)
}
