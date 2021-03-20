package parser

type Edge struct {
	LeftColumn             int
	RightColumn            int
	TopLine                int
	BottomLine             int
	Label                  string
	DestinationLeftOrUp    *Node
	DestinationRightOrDown *Node
}
