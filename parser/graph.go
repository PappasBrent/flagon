package parser

type Graph struct {
	Nodes        []*Node
	Edges        []*Edge
	LabeledNodes map[string]*Node
	LabeledEdges map[string]*Edge
}
