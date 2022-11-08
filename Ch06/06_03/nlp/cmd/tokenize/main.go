package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/353solutions/nlp"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: tokenize SENTENCE")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	for _, tok := range nlp.Tokenize(flag.Arg(0)) {
		fmt.Println(tok)
	}
}
