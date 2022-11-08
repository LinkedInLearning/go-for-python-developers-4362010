package main

import (
	"fmt"
	"os"
)

func fileHead(fileName string, size int) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, size)
	n, err := file.Read(buf)

	if err != nil {
		return nil, err
	}
	if n != size {
		return nil, fmt.Errorf("%q too small", fileName)
	}

	return buf, nil
}

func main() {
	data, err := fileHead("head.png", 8)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(data)
	}
}
