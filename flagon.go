package flagon

import "github.com/PappasBrent/flagon/parser"

func Parse(text string) (*parser.Graph, error) {
	return parser.Parse(text)
}
