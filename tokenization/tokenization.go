package tokenization

import "unicode"

func Tokenize(text string) ([]*Token, *TokenError) {
	tokens := make([]*Token, 0)
	line := 1
	column := 0
	i := -1
	for i < len(text) {
		i++

		// Handle whitespace
		if i >= len(text) {
			break
		}
		column++
		if text[i] == '\n' {
			line++
			column = 0
			continue
		} else if unicode.IsSpace(rune(text[i])) {
			continue
		}

		// Lex a token
		token := Token{Line: line, Column: column}
		if text[i] == '[' {
			token.Type = LeftBracket
			token.Value = "["
		} else if text[i] == ']' {
			token.Type = RightBracket
			token.Value = "]"
		} else if text[i] == '-' {
			token.Type = Dash
			token.Value = "-"
		} else if text[i] == '|' {
			token.Type = Pipe
			token.Value = "|"
		} else if unicode.IsLetter(rune(text[i])) {
			token.Type = Label
			for (unicode.IsLetter(rune(text[i])) || text[i] == ' ') && i < len(text) {
				token.Value += string(text[i])
				i++
				column++
			}
			// Unread last character
			i--
			column--
		} else {
			return nil, &TokenError{line, column, rune(text[i])}
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}
