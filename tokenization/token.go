package tokenization

// Token reprents a token recognized by the lexer
type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}
