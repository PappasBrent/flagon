package parser

type Node struct {
	Line        int
	LeftColumn  int
	RightColumn int
	Label       string
	EdgeLeft    *Edge
	EdgeRight   *Edge
	EdgeUp      *Edge
	EdgeDown    *Edge
}
