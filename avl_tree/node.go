package avl

type Node struct {
	hash  []byte
	value []byte
	left  *Node
	right *Node
}
