package binarysearchtree

import (
	"testing"
)

func constructTree() *TreeNode {
	root := TreeNode{5, nil, nil}
	a := TreeNode{3, nil, nil}
	b := TreeNode{6, nil, nil}
	c := TreeNode{2, nil, nil}
	d := TreeNode{4, nil, nil}
	e := TreeNode{7, nil, nil}
	root.Left = &a
	root.Right = &b
	a.Left = &c
	a.Right = &d
	b.Right = &e
	return &root
}

func TestInsertNode(t *testing.T) {
	traverseTree(insertNode(constructTree(), 0))
}

func TestTraverseTree(t *testing.T) {
	traverseTree(constructTree())
}
