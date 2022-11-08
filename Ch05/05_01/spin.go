package main

import (
	"fmt"
	"time"
)

func worker(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d\n", n)
}

func main() {
	for i := 0; i < 5; i++ {
		go worker(i)
	}
	fmt.Println("main")
}
