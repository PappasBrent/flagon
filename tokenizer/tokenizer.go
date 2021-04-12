package tokenizer

import (
	"fmt"
	"unicode"
)

// Token reprents a token recognized by the lexer
type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

type TokenType string

const (
	LeftBracket  TokenType = "LeftBracket"
	RightBracket TokenType = "RightBracket"
	Dash         TokenType = "Dash"
	Pipe         TokenType = "Pipe"
	Label        TokenType = "Label"
)

type TokenError struct {
	Line   int
	Column int
	Value  rune
}

func (t *TokenError) Error() string {
	return fmt.Sprintf("Unexpected character at %d %d: %q",
		t.Line, t.Column, t.Value)
}

type StopIteration struct{}

func (s *StopIteration) Error() string {
	return "Input exhausted"
}

type Tokenizer struct {
	Input         string
	CurrentRune   rune
	NextRune      rune
	CurrentLine   int
	CurrentColumn int
	index         int
}

func New(input string) Tokenizer {
	return Tokenizer{input, 0, 0, 1, 0, -1}
}

// TokenIterator is the interface for Tokenizers.
//
// Next returns the next token in the input string
// and any error encountered during lexing
type TokenIterator interface {
	Next() (token Token, err error)
}

// advance advances the tokenizer's current position in the input string.
// It returns a StopIteration error if the end of the input has been reached.
func (tokenizer *Tokenizer) advance() error {
	tokenizer.index++
	if tokenizer.index >= len(tokenizer.Input) {
		return &StopIteration{}
	}
	tokenizer.CurrentColumn++
	tokenizer.CurrentRune = rune(tokenizer.Input[tokenizer.index])
	tokenizer.NextRune = 0
	if tokenizer.index+1 < len(tokenizer.Input) {
		tokenizer.NextRune = rune(tokenizer.Input[tokenizer.index+1])
	}
	return nil
}

// Next implements the TokenIterator interface:
// it returns the next token in the input string,
// and any error encountered during lexing
func (tokenizer *Tokenizer) Next() (Token, error) {
	err := tokenizer.advance()
	if err != nil {
		return Token{}, err
	}
	for tokenizer.CurrentRune == rune('\n') {
		tokenizer.CurrentLine++
		tokenizer.CurrentColumn = 0
		err := tokenizer.advance()
		if err != nil {
			return Token{}, err
		}
	}
	for unicode.IsSpace(tokenizer.CurrentRune) {
		err := tokenizer.advance()
		if err != nil {
			return Token{}, err
		}
	}
	token := Token{Line: tokenizer.CurrentLine, Column: tokenizer.CurrentColumn}
	if tokenizer.CurrentRune == rune('[') {
		token.Type = LeftBracket
		token.Value = "["
	} else if tokenizer.CurrentRune == rune(']') {
		token.Type = RightBracket
		token.Value = "]"
	} else if tokenizer.CurrentRune == rune('-') {
		token.Type = Dash
		token.Value = "-"
	} else if tokenizer.CurrentRune == '|' {
		token.Type = Pipe
		token.Value = "|"
	} else if unicode.IsLetter(tokenizer.CurrentRune) {
		token.Type = Label
		for unicode.IsLetter(tokenizer.CurrentRune) || tokenizer.CurrentRune == rune(' ') {
			token.Value += string(tokenizer.CurrentRune)
			if !unicode.IsLetter(tokenizer.NextRune) && !(tokenizer.NextRune == rune(' ')) {
				break
			}
			err = tokenizer.advance()
			if err != nil {
				return Token{}, err
			}
		}
	}
	return token, nil
}
