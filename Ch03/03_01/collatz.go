package main

import "fmt"

// CollatzStep computes one step in CollatzStep conjecture
func CollatzStep(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func main() {
	fmt.Println(CollatzStep(4))
	fmt.Println(CollatzStep(5))
}
