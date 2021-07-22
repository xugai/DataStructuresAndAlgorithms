package treetraverse

import (
	"testing"
)

func constructTree() *Tree {
	root := Tree{"A", nil, nil}
	B := Tree{"B", nil, nil}
	C := Tree{"C", nil, nil}
	D := Tree{"D", nil, nil}
	E := Tree{"E", nil, nil}
	F := Tree{"F", nil, nil}
	G := Tree{"G", nil, nil}
	B.left = &D
	B.right = &E
	C.left = &F
	C.right = &G
	root.left = &B
	root.right = &C
	return &root
}

func TestPreOrder(t *testing.T) {
	root := constructTree()
	PreOrder(root)
}

func TestInOrder(t *testing.T) {
	root := constructTree()
	InOrder(root)
}

func TestPostOrder(t *testing.T) {
	root := constructTree()
	PostOrder(root)
}

func TestSequentialOrder(t *testing.T) {
	root := constructTree()
	SequentialOrder(root)
}

//  A B C D E F G
//  1 2 3 4 5 6 7
