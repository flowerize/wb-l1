package main

import "fmt"

type OldCalculator interface {
	Operation(num1, num2 int, op string) int
}

type NewCalculator struct{}

func (nc NewCalculator) Add(a, b int) int  { return a + b }
func (nc NewCalculator) Mult(a, b int) int { return a * b }
func (nc NewCalculator) Div(a, b int) int  { return a / b }

// Адаптер, реализующий старый интерфейс
type CalculatorAdapter struct {
	calculator NewCalculator
}

func (ca CalculatorAdapter) Operation(num1, num2 int, op string) int {
	switch op {
	case "add":
		return ca.calculator.Add(num1, num2)
	case "mult":
		return ca.calculator.Mult(num1, num2)
	case "div":
		return ca.calculator.Div(num1, num2)
	default:
		return 0
	}
}

func main() {
	adapter := CalculatorAdapter{calculator: NewCalculator{}}

	fmt.Println(adapter.Operation(2, 3, "add"))
	fmt.Println(adapter.Operation(4, 5, "mult"))
	fmt.Println(adapter.Operation(10, 2, "div"))
}
