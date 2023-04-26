package avl

import "fmt"

type Tree struct {
	root *Node
}

func NewAVLTree(root *Node) *Tree {
	return &Tree{
		root: root,
	}
}

func (t *Tree) Set(key, value []byte) error {
	if key == nil {
		return fmt.Errorf("error while set: cannot set empty key")
	}

	if t.root == nil {
		t.root = NewNode(key, value, 0, 1)
		return nil
	}

	return t.root.set(key, value)
}

func (t *Tree) Delete(key []byte) ([]byte, error) {
	// no-op: implement
	return nil, nil
}

func (t *Tree) Get(key []byte) ([]byte, error) {
	// no-op: implement
	return nil, nil
}

func (t Tree) Height() uint64 {
	return t.root.height
}
