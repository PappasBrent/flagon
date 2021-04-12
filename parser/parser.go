package parser

import (
	"fmt"

	"github.com/PappasBrent/flagon/tokenizer"
)

type graphParser struct {
	CurrentTokenIndex int
	Tokens            []tokenizer.Token
	CurrentToken      tokenizer.Token
	UnusedTokens      map[tokenizer.Token]bool
	TokenMap          map[int]map[int]tokenizer.Token
	NodeMap           map[int]map[int]*Node
	EdgeHorizontalMap map[int]map[int]*Edge
	EdgeVerticalMap   map[int]map[int]*Edge
	Graph             *Graph
}

func newGraphParser(tokens []tokenizer.Token) *graphParser {
	tokenMap := make(map[int]map[int]tokenizer.Token)
	unusedTokens := make(map[tokenizer.Token]bool)
	for _, token := range tokens {
		if tokenMap[token.Line] == nil {
			tokenMap[token.Line] = make(map[int]tokenizer.Token)
		}
		tokenMap[token.Line][token.Column] = token
		unusedTokens[token] = true
	}
	p := graphParser{
		-1,
		tokens,
		tokenizer.Token{},
		unusedTokens,
		tokenMap,
		make(map[int]map[int]*Node),
		make(map[int]map[int]*Edge),
		make(map[int]map[int]*Edge),
		&Graph{
			make([]*Node, 0),
			make([]*Edge, 0),
			make(map[string]*Node),
			make(map[string]*Edge),
		},
	}
	for _, t := range p.Tokens {
		fmt.Println(t)
	}
	p.advanceForward()
	return &p
}

// Gets the next token to the right based upon the current token location
func (p *graphParser) getNextTokenRight() tokenizer.Token {
	if p.CurrentToken == (tokenizer.Token{}) {
		return p.Tokens[0]
	}
	for _, token := range p.Tokens[p.CurrentTokenIndex+1:] {
		if token.Line == p.CurrentToken.Line && token.Column > p.CurrentToken.Column {
			return token
		}
	}
	return tokenizer.Token{}
}

// Gets the next token down based upon the current token location
func (p *graphParser) getNextTokendown() tokenizer.Token {
	if p.CurrentToken == (tokenizer.Token{}) {
		panic("Current token may not be null when getting next token down")
	}
	for _, token := range p.Tokens[p.CurrentTokenIndex+1:] {
		if token.Line > p.CurrentToken.Line && token.Column == p.CurrentToken.Column {
			return token
		}
	}
	return tokenizer.Token{}
}

// Advance to the next token in the stream
func (p *graphParser) advanceForward() {
	p.CurrentTokenIndex++
	if p.CurrentTokenIndex >= len(p.Tokens) {
		p.CurrentToken = tokenizer.Token{}
		return
	}
	delete(p.UnusedTokens, p.CurrentToken)
	p.CurrentToken = p.Tokens[p.CurrentTokenIndex]
}

// Advances the parser's current token to the right
func (p *graphParser) advanceRight() {
	delete(p.UnusedTokens, p.CurrentToken)
	p.CurrentToken = p.getNextTokenRight()
}

// Advances the parser's current token down
func (p *graphParser) advanceDown() {
	delete(p.UnusedTokens, p.CurrentToken)
	p.CurrentToken = p.getNextTokendown()
}

func (p *graphParser) acceptRight(expected tokenizer.TokenType) bool {
	nextTokenRight := p.getNextTokenRight()
	if nextTokenRight == (tokenizer.Token{}) || nextTokenRight.Type != expected {
		return false
	} else {
		p.advanceRight()
		return true
	}
}

func (p *graphParser) acceptDown(expected tokenizer.TokenType) bool {
	nextTokenDown := p.getNextTokendown()
	if nextTokenDown == (tokenizer.Token{}) || nextTokenDown.Type != expected {
		return false
	} else {
		p.advanceDown()
		return true
	}
}

func (p *graphParser) parseNode() *Node {
	// Current token: LeftBracket
	// Node -> LeftBracket [Label] RightBracket
	// Should make a new function, ParseItems
	node := &Node{
		p.CurrentToken.Line,
		p.CurrentToken.Column,
		-1,
		"",
		nil,
		nil,
		nil,
		nil,
	}
	if p.acceptRight(tokenizer.Label) {
		node.Label = p.CurrentToken.Value
	}
	if !p.acceptRight(tokenizer.RightBracket) {
		s := fmt.Sprintf("Expected Rightbracket to close node %d %d", p.CurrentToken.Line, p.CurrentToken.Column)
		panic(s)
	}
	node.RightColumn = p.CurrentToken.Column
	return node
}

func (p *graphParser) parseEdgeHorizontal() *Edge {
	// Current token: Dash
	// HorizontalEdge -> Dash {Dash} [Label Dash {Dash}]
	edge := &Edge{
		p.CurrentToken.Column,
		-1,
		p.CurrentToken.Line,
		p.CurrentToken.Line,
		"",
		nil,
		nil,
	}
	for p.acceptRight(tokenizer.Dash) {
	}
	if p.acceptRight(tokenizer.Label) {
		edge.Label = p.CurrentToken.Value
		if !p.acceptRight(tokenizer.Dash) {
			panic("Expected a dash after lock name")
		}
		for p.acceptRight(tokenizer.Dash) {
		}
	}
	edge.RightColumn = p.CurrentToken.Column
	return edge
}

func (p *graphParser) parseEdgeVertical() *Edge {
	// Current token: Pipe
	// VerticalEdge (parses down) -> Pipe {Pipe} [Label Pipe {Pipe}]
	edge := &Edge{
		p.CurrentToken.Column,
		p.CurrentToken.Column,
		p.CurrentToken.Line,
		-1,
		"",
		nil,
		nil,
	}
	for p.acceptDown(tokenizer.Pipe) {
	}
	if p.acceptDown(tokenizer.Label) {
		edge.Label = p.CurrentToken.Value
		if !p.acceptDown(tokenizer.Pipe) {
			panic("Expected a pipe after lock name")
		}
		for p.acceptDown(tokenizer.Pipe) {
		}
	}
	edge.BottomLine = p.CurrentToken.Line
	return edge
}

func (p *graphParser) parseNodesAndEdges() {
	for p.CurrentToken != (tokenizer.Token{}) {
		if _, unused := p.UnusedTokens[p.CurrentToken]; !unused {
			p.advanceForward()
			continue
		}

		if p.CurrentToken.Type == tokenizer.LeftBracket {
			node := p.parseNode()
			if p.NodeMap[node.Line] == nil {
				p.NodeMap[node.Line] = make(map[int]*Node)
			}
			p.NodeMap[node.Line][node.LeftColumn] = node

			// Add node to graph
			if node.Label != "" {
				if _, exists := p.Graph.LabeledNodes[node.Label]; exists {
					panic("Error: Two nodes with label " + node.Label)
				}
				p.Graph.LabeledNodes[node.Label] = node
			}
			p.Graph.Nodes = append(p.Graph.Nodes, node)
		} else if p.CurrentToken.Type == tokenizer.Dash {
			edge := p.parseEdgeHorizontal()
			if p.EdgeHorizontalMap[edge.TopLine] == nil {
				p.EdgeHorizontalMap[edge.TopLine] = make(map[int]*Edge)
			}
			p.EdgeHorizontalMap[edge.TopLine][edge.LeftColumn] = edge
			p.EdgeHorizontalMap[edge.TopLine][edge.RightColumn] = edge

			p.Graph.Edges = append(p.Graph.Edges, edge)
		} else if p.CurrentToken.Type == tokenizer.Pipe {
			edge := p.parseEdgeVertical()
			if p.EdgeVerticalMap[edge.TopLine] == nil {
				p.EdgeVerticalMap[edge.TopLine] = make(map[int]*Edge)
			}
			p.EdgeVerticalMap[edge.TopLine][edge.LeftColumn] = edge
			if p.EdgeVerticalMap[edge.BottomLine] == nil {
				p.EdgeVerticalMap[edge.BottomLine] = make(map[int]*Edge)
			}
			p.EdgeVerticalMap[edge.BottomLine][edge.LeftColumn] = edge

			p.Graph.Edges = append(p.Graph.Edges, edge)
		} else {
			panic(
				fmt.Sprintf("Invalid architecture token found: %q", p.CurrentToken),
			)
		}

		p.advanceForward()
	}
}

func (p *graphParser) connectNodesAndEdges() {
	// Link nodes to edges
	for _, nodesInLine := range p.NodeMap {
		for _, node := range nodesInLine {
			// If checks for linking edges to nodes
			node.EdgeLeft = p.EdgeHorizontalMap[node.Line][node.LeftColumn-1]
			if node.EdgeLeft != nil {
				node.EdgeLeft.DestinationRightOrDown = node
			}
			node.EdgeRight = p.EdgeHorizontalMap[node.Line][node.RightColumn+1]
			if node.EdgeRight != nil {
				node.EdgeRight.DestinationLeftOrUp = node
			}
			// Vertical edges must be aligned to nodes' left columns
			node.EdgeUp = p.EdgeVerticalMap[node.Line-1][node.LeftColumn]
			if node.EdgeUp != nil {
				node.EdgeUp.DestinationRightOrDown = node
			}
			node.EdgeDown = p.EdgeVerticalMap[node.Line+1][node.LeftColumn]
			if node.EdgeDown != nil {
				node.EdgeDown.DestinationLeftOrUp = node
			}

		}
	}

	edgeMaps := []map[int]map[int]*Edge{p.EdgeHorizontalMap, p.EdgeVerticalMap}

	for _, edgeMap := range edgeMaps {
		for _, edgesInLine := range edgeMap {
			for _, edge := range edgesInLine {
				// Check if an edge has no nodes connected to it
				if edge.DestinationLeftOrUp == nil && edge.DestinationRightOrDown == nil {
					panic("Warning: Edge to no nodes found")
				}

				// Add edge to graph
				if edge.Label != "" {
					if foundEdge, exists := p.Graph.LabeledEdges[edge.Label]; exists && foundEdge != edge {
						panic("Error: Two edges with label " + edge.Label)
					}
					p.Graph.LabeledEdges[edge.Label] = edge
				}
			}
		}
	}
}

// Parses a given piece of text into a graph
func Parse(text string) (*Graph, error) {
	tzr := tokenizer.New(text)
	tokens := make([]tokenizer.Token, 0)
	var token tokenizer.Token
	var err error
	for token, err = tzr.Next(); err == nil; token, err = tzr.Next() {
		tokens = append(tokens, token)
	}
	if _, stopped := err.(*tokenizer.StopIteration); !stopped {
		return nil, err
	}

	p := newGraphParser(tokens)
	p.parseNodesAndEdges()
	p.connectNodesAndEdges()

	return p.Graph, nil
}
