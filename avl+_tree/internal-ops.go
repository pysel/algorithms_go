package avl

import (
	"bytes"
	"fmt"
	"math"
)

func (n *Node) set(key, value []byte) error {
	if n.isLeaf() {
		newNode := NewNode(key, value, 0, 1)
		switch bytes.Compare(key, n.key) {
		case -1: // means key we are trying to insert is less than the key of the current node, node has to be added/modified in the left subtree
			n.left = newNode
			if n.right == nil { // increase the height of the recursed by node by 1 if we just added the first node to it's children
				n.height++
			}
		case 1: // means key we are trying to insert is greater than the key of the current node, node has to be added/modified in the right subtree
			n.right = newNode
			if n.left == nil {
				n.height++
			}
		case 0: // means key is equal to the one we are trying to insert
			n = newNode
		}

		return nil
	}

	// need to check if we need to update the height of the current Node
	rightHeight := n.right.height
	leftHeight := n.left.height

	var err error
	if bytes.Compare(key, n.key) < 0 {
		err = n.left.set(key, value)
	} else {
		err = n.right.set(key, value)
	}

	if err != nil {
		return fmt.Errorf("error while setting a value %s", err)
	}

	// if maximum sub height is different from the maximum recorded height, we need to update current Node's height
	if returnMaxSubHeight(n) > maxUint64(rightHeight, leftHeight) {
		n.height++
	}

	return nil
}

func returnMaxSubHeight(n *Node) uint64 {
	return uint64(math.Max(float64(n.right.height), float64(n.left.height)))
}

func maxUint64(a, b uint64) uint64 {
	return uint64(math.Max(float64(a), float64(b)))
}
