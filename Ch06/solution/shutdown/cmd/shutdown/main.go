package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/353solutions/shutdown"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: shutdown PID_FILE\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	if err := shutdown.ShutdownServer(flag.Arg(0)); err != nil {
		os.Exit(1)
	}
}
