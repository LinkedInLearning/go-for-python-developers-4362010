package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		text     string
		expected []string
	}{
		{"", nil},
		{"hi", []string{"hi"}},
		{"HI", []string{"hi"}},
		{"Who's on first.", []string{"who", "on", "first"}},
	}

	for _, tc := range testCases {
		name := tc.text
		if len(name) == 0 {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			out := Tokenize(tc.text)
			require.Equal(tc.expected, out)
		})
	}
}
