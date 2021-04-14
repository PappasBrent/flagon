Flagon
=============================================
[![Build Status](https://travis-ci.com/PappasBrent/flagon.svg?branch=main)](https://travis-ci.com/PappasBrent/flagon)
[![codecov](https://codecov.io/gh/PappasBrent/flagon/branch/main/graph/badge.svg?token=OQPCDHSA95)](https://codecov.io/gh/PappasBrent/flagon)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
Run this command from your Go project directory to add Flagon to your project's
go.mod file:
```
go get github.com/PappasBrent/flagon
```

## Quick Start

Import Flagon:
```go
import "github.com/PappasBrent/flagon"
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

Parse the string variable with a call to the `Parse` function:
```go
graph, _ := flagon.Parse(text)
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

graph, _ := flagon.Parse(text)
```

### Getting Nodes
Graph structs have a LabeledNodes field containing a mapping of all labels to their
corresponding Node structs. Labels are strings.

Example:
```go
  text := `[A]-AB-[B]
           |
           |
           []
`

    graph, _ := flagon.Parse(text)
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


All of a graph's nodes are stored in its Nodes field.
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
Graph structs have a LabeledEdges field containing a mapping of all labels to their
corresponding Edge structs. Labels are strings.

Example:
```go
  for label, edge := range graph.LabeledEdges {
    fmt.Printf("Parsed an edge with label %v with top-left at"+
      " %v:%v and bottom-right at %v:%v\n", label, edge.TopLine,
      edge.LeftColumn, edge.BottomLine, edge.RightColumn)
  }
```

Output:
```
Parsed an edge with label AB with top-left at 1:4 and bottom-right at 1:7
```


All of a graph's edges are stored in its Edges field.
Continuing from the previous example:
```go
  for _, edge := range graph.Edges {
    fmt.Printf("Parsed an edge with top-left at"+
      " %v:%v and bottom-right at %v:%v\n", edge.TopLine,
      edge.LeftColumn, edge.BottomLine, edge.RightColumn)
  }
```

Output:
```
Parsed an edge with top-left at 1:4 and bottom-right at 1:7
Parsed an edge with top-left at 2:14 and bottom-right at 3:14
```

### Traversing a Graph

## Testing
Use `go test` to run Flagon's tests.

To run the tokenization tests use `go test ./tokenization`

Parser tests have not yet been implemented. Use `go test ./parser` to run
the parser tests once they have been added.

## Acknowledgments
- The [Funciton esoteric programming language](https://esolangs.org/wiki/Funciton)
for the idea
- This research paper for suggestions on the implementation:
Tomita M. (1991) Parsing 2-Dimensional Language. In: Tomita M. (eds) Current Issues in Parsing Technology. The Springer International Series in Engineering and Computer Science (Natural Language Processing and Machine Translation), vol 126. Springer, Boston, MA. https://doi.org/10.1007/978-1-4615-3986-5_18
