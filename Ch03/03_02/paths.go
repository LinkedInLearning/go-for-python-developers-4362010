package main

import (
	"fmt"
	"strings"
)

func SplitExt(path string) (string, string) {
	i := strings.LastIndex(path, ".")
	if i == -1 {
		return path, ""
	}

	return path[:i], path[i:]
}

func main() {
	root, ext := SplitExt("app.go")
	fmt.Printf("root=%#v, path=%#v\n", root, ext)
}
