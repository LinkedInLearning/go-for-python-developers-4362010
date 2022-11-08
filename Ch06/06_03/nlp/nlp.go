package nlp

import (
	"regexp"
	"strings"

	"github.com/353solutions/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile("[a-z]+")
)

// Tokenize splits text to list of tokens.
func Tokenize(text string) []string {
	text = strings.ToLower(text)
	var tokens []string
	for _, tok := range wordRe.FindAllString(text, -1) {
		tok = stemmer.Stem(tok)
		if tok == "" {
			continue
		}
		tokens = append(tokens, tok)
	}
	return tokens
}
