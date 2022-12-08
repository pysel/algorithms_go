package avl

// Node is a node in the tree
type Node struct {
	hash  []byte
	value []byte
	left  *Node
	right *Node
}
