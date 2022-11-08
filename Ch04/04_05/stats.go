package main

import "fmt"

func Add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.0, 2.0))
	fmt.Println(Add("G", "o"))
}
