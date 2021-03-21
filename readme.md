Flagon
=============================================

<img src="./logo.png" height="250" align="right"/>
An ASCII graph parser written in Go.
Flagon stands for FLuid ASCII Graph Object Notation.

## Table of Contents
- [Flagon](#flagon)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
  - [Examples](#examples)
    - [Parsing an ASCII Graph](#parsing-an-ascii-graph)
    - [Getting Nodes](#getting-nodes)
    - [Getting Edges](#getting-edges)
    - [Traversing a Graph](#traversing-a-graph)
  - [Testing](#testing)
  - [Acknowledgments](#acknowledgments)

## Installation

## Quick Start

Import the flagon parser:
```go
import "github.com/PappasBrent/flagon/parser"
```

Assign a string variable to the ASCII graph you would like to Flagon to parse:
```go
text := `[A]-AB-[B]
         |
         AC
         |
         [C]
`
```

Parse the string variable with a call to the parser package's `Parse` method:
```go
graph, _ := parser.Parse(text)
```

Unnecessary labels for nodes and edges can be omitted, e.g., this is valid:
```go
text := `[A]-AB-[B]
         |
         |
         []
`
```

## Examples

### Parsing an ASCII Graph
```go
text := `[A]-AB-[B]
         |
         |
         []
`

graph, _ := parser.Parse(text)
```

### Getting Nodes
Graph structs have a LabeledNodes member containing a mapping of all labels to their
corresponding Node structs. Labels are strings.

Example:
```go
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
```

Output:
```
Parsed a node with label A on line 1 with left bracket at column 1
Parsed a node with label B on line 1 with left bracket at column 8
```


All of a graph's nodes are stored in its nodes member.
Continuing from the previous example:
```go
  for _, node := range graph.Nodes {
	  fmt.Printf("Parsed a node starting at line %v column %v\n",
        node.Line, node.LeftColumn)
	}
```

Output:
```
Parsed a node starting at line 1 column 1
Parsed a node starting at line 1 column 8
Parsed a node starting at line 4 column 14
```

### Getting Edges
### Traversing a Graph

## Testing

## Acknowledgments