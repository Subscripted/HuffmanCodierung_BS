package types

type Node struct {
	Char        rune
	Left, Right *Node
	Frequency   int
}
