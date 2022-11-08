package main

import (
	"fmt"
	"os"
)

type response struct {
	fileName string
	size     int64
}

func worker(fileName string, ch chan<- response) {
	resp := response{fileName: fileName, size: -1}
	st, err := os.Stat(fileName)
	if err == nil {
		resp.size = st.Size()
	}
	ch <- resp
}

func main() {

	files := []string{
		"files/a.txt",
		"files/b.txt",
		"files/c.txt",
		"files/d.txt",
		"files/e.txt",
	}
	ch := make(chan response)

	for _, f := range files {
		go worker(f, ch)
	}

	for range files {
		r := <-ch
		fmt.Printf("%s -> %d\n", r.fileName, r.size)
	}
}
