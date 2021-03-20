package tokenization

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected []*Token
	}{
		{
			"Tokenize an unlabeled node",
			"[]",
			[]*Token{
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
			[]*Token{
				{LeftBracket, "[", 2, 13},
				{RightBracket, "]", 2, 14},
				{Pipe, "|", 3, 13},
				{Pipe, "|", 4, 13},
			},
		},
	}
	for i, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := Tokenize(tC.input)
			if err != nil {
				t.Errorf("Error running test %d", i+1)
			}
			if !reflect.DeepEqual(result, tC.expected) {
				t.Logf("Test %d failed\nExpected: ", i+1)
				for _, token := range tC.expected {
					t.Logf("%v", token)
				}
				t.Log("Got: ")
				for _, token := range result {
					t.Logf("%v", token)
				}
				t.Fatal("Failed")
			}
		})
	}
}
