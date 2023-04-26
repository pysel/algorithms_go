package avl

import (
	"fmt"
	"testing"
)

var (
	defaultSets = []KV{
		{
			key:   []byte("hello"),
			value: []byte("world"),
		},
		{
			key:   []byte("hello1"),
			value: []byte("world"),
		},
		{
			key:   []byte("hello2"),
			value: []byte("world"),
		},
		{
			key:   []byte("hello3"),
			value: []byte("world"),
		},
	}
)

type (
	KV struct {
		key   []byte
		value []byte

		// TODO: uncomment
		// delete bool
	}
)

func TestSet(t *testing.T) {
	tree := NewAVLTree(nil)

	// invalid: trying to set empty key
	err := tree.Set(nil, defaultSets[0].value)
	fmt.Println(err)
	// require.Error(t, err)

	err = tree.Set(defaultSets[0].key, defaultSets[0].value)
	// require.NoError(t, err)

	err = tree.Set(defaultSets[1].key, defaultSets[0].value)

	fmt.Println(string(tree.root.value))
}
