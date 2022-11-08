/*
Print character and frequency in percent of Rumi's poem
Print in frequency descending order
Assume ASCII text, ignore white space
*/
package main

import (
	"fmt"
	"sort"
	"unicode"
)

var poem = `
those who do not feel this love
pulling them like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let them sleep
`

func main() {
	counts := make(map[rune]int)
	count := 0.0
	for _, c := range poem {
		if unicode.IsSpace(c) {
			continue
		}
		counts[c]++
		count += 1
	}

	var chars []rune
	for c := range counts {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool {
		c1, c2 := chars[i], chars[j]
		return counts[c1] > counts[c2]
	})

	for _, c := range chars {
		n := counts[c]
		f := float64(n) / count * 100
		fmt.Printf("%c: %.2f\n", c, f)
	}
}
