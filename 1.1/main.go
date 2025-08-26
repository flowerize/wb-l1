package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Weight float64
	Height float64
	Rating float64
}

func (h Human) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d y.o.\n", h.Name, h.Age)
}

func (h Human) Parametres() {
	fmt.Printf("I have %.1f kg and %.1f cm\n", h.Weight, h.Height)
}

func (h *Human) Ranking(r float64) {
	h.Rating = r
	var tier string
	switch {
	case r >= 0 && r <= 1.5:
		tier = "tier C"
	case r > 1.5 && r <= 2.9:
		tier = "B-Tier"
	case r > 2.9 && r <= 4.0:
		tier = "A-Tier"
	case r > 4.0 && r <= 5.0:
		tier = "S-Tier"
	default:
		tier = "Unknown"
	}
	fmt.Printf("Rating: %.1f - %s\n", r, tier)
}

type Action struct {
	Human // Встраивание (embedding)
	Task  string
}

func main() {
	a := Action{
		Human: Human{
			Name:   "Alice",
			Age:    21,
			Weight: 56.36,
			Height: 172.51},
		Task: "Coding",
	}

	a.Greet()
	a.Parametres()
	a.Ranking(3.1)

	fmt.Println("Task:", a.Task)
}
