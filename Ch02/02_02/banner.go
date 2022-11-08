/*
Print a banner with length of 8 for the string "Hi ☺"
Should output

	  Hi ☺
	--------
*/
package main

import "fmt"

func main() {
	message := "Hi ☺"
	width := 8
	padding := (width - len(message)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(message)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
