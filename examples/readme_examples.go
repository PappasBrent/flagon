package examples

import (
	"fmt"

	"github.com/PappasBrent/flagon/parser"
)

func Parse_labeled_graph() {
	text := `[A]-AB-[B]
             |
             |
             []
`

	graph, _ := parser.Parse(text)
	for label, node := range graph.LabeledNodes {
		fmt.Printf("Parsed a node with label %v on line %v"+
			" with left bracket at column %v\n",
			label, node.Line, node.LeftColumn)
	}

	for _, node := range graph.Nodes {
		fmt.Printf("Parsed a node starting at line %v column %v\n",
			node.Line, node.LeftColumn)
	}

	fmt.Println()

	for label, edge := range graph.LabeledEdges {
		fmt.Printf("Parsed an edge with label %v with top-left at"+
			" %v:%v and bottom-right at %v:%v\n", label, edge.TopLine,
			edge.LeftColumn, edge.BottomLine, edge.RightColumn)
	}

	for _, edge := range graph.Edges {
		fmt.Printf("Parsed an edge with top-left at"+
			" %v:%v and bottom-right at %v:%v\n", edge.TopLine,
			edge.LeftColumn, edge.BottomLine, edge.RightColumn)
	}
}
