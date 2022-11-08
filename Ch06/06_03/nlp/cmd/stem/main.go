package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/353solutions/nlp/stemmer"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: stem WORD")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	fmt.Println(stemmer.Stem(flag.Arg(0)))
}
