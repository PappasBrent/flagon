package tokenizer

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected []Token
	}{
		{
			"Tokenize an unlabeled node",
			"[]",
			[]Token{
				{LeftBracket, "[", 1, 1},
				{RightBracket, "]", 1, 2},
			},
		},
		{
			"Tokenize over multiple lines",
			`
            []
            |
            |`,
			[]Token{
				{LeftBracket, "[", 2, 13},
				{RightBracket, "]", 2, 14},
				{Pipe, "|", 3, 13},
				{Pipe, "|", 4, 13},
			},
		},
		{
			"Tokenize labeled nodes and edges",
			`
            [A]--AB--[B]
            |
            |
            AC
            |
            |
            [C]`,
			[]Token{
				{LeftBracket, "[", 2, 13},
				{Label, "A", 2, 14},
				{RightBracket, "]", 2, 15},
				{Dash, "-", 2, 16},
				{Dash, "-", 2, 17},
				{Label, "AB", 2, 18},
				{Dash, "-", 2, 20},
				{Dash, "-", 2, 21},
				{LeftBracket, "[", 2, 22},
				{Label, "B", 2, 23},
				{RightBracket, "]", 2, 24},
				{Pipe, "|", 3, 13},
				{Pipe, "|", 4, 13},
				{Label, "AC", 5, 13},
				{Pipe, "|", 6, 13},
				{Pipe, "|", 7, 13},
				{LeftBracket, "[", 8, 13},
				{Label, "C", 8, 14},
				{RightBracket, "]", 8, 15},
			},
		},
	}
	for i, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := make([]Token, 0)
			tokenizer := New(tC.input)
			var token Token
			var err error
			for token, err = tokenizer.Next(); err == nil; token, err = tokenizer.Next() {
				result = append(result, token)
			}
			if _, stopped := err.(*StopIteration); !stopped {
				t.Errorf("Test %q\n%v", tC.desc, err)
			}
			if !reflect.DeepEqual(result, tC.expected) {
				t.Logf("Test %d failed\nOutput: ", i+1)
				for _, token := range result {
					t.Logf("%v", token)
				}
				t.Log("\nExpected: ")
				for _, token := range tC.expected {
					t.Logf("%v", token)
				}
				t.Fatal("Failed")
			}
		})
	}
}
