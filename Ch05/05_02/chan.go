package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 99    // send
	val := <-ch // receive
	fmt.Printf("got %d\n", val)
}
