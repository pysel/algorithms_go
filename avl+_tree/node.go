package avl

import "math"

// Only leaf values have values
// Internal nodes have key = key of leftmost node in the right branch
type Node struct {
	hash   []byte
	key    []byte
	value  []byte
	height uint64
	size   uint64
	left   *Node
	right  *Node
}

func NewNode(key, value []byte, height, size uint64) *Node {
	return &Node{
		key:    key,
		value:  value,
		height: height,
		size:   size,
		left:   nil,
		right:  nil,
		hash:   nil,
	}
}

func (n *Node) Balance() int64 {
	diff := float64(n.left.height - n.right.height)
	return int64(math.Abs(diff))
}

func (n *Node) isLeaf() bool {
	return n.height == 0
}
