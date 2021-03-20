package tokenization

type TokenType string

const (
	LeftBracket  TokenType = "LeftBracket"
	RightBracket TokenType = "RightBracket"
	Dash         TokenType = "Dash"
	Pipe         TokenType = "Pipe"
	Label        TokenType = "Label"
)
