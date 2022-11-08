package nlp_test

import (
	"fmt"

	"github.com/353solutions/nlp"
)

func ExampleTokenize() {
	s := "Hi, how are you feeling today?"
	tokens := nlp.Tokenize(s)
	fmt.Println(tokens)
	// Output:
	// [hi how are you feel today]
}
