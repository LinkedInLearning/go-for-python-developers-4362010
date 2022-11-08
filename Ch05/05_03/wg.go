package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for msg := range ch {
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%d finished %s\n", id, msg)
				wg.Done()
			}
		}(i)
	}

	for _, msg := range []string{"A", "B", "C", "D", "E", "F"} {
		wg.Add(1)
		ch <- msg
	}
	wg.Wait()
	fmt.Println("all jobs done")
}
