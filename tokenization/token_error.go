package tokenization

import "fmt"

type TokenError struct {
	Line   int
	Column int
	Value  rune
}

func (t *TokenError) Error() string {
	return fmt.Sprintf("Unexpected character at %d %d: %q",
		t.Line, t.Column, t.Value)
}
