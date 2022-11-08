// What is the most common words in Rumi's poem?
package main

import (
	"strings"
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
	frequency := make(map[string]int)
	for _, word := range strings.Fields(poem) {
		frequency[word]++
	}

	maxW, maxC := "", 0
	for w, c := range frequency {
		if c > maxC {
			maxW, maxC = w, c
		}
	}

	print(maxW)
}
