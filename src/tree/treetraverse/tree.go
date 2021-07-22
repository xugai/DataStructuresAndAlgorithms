package treetraverse

import (
	"fmt"
	"math"
)

type Tree struct {
	val string
	left *Tree
	right *Tree
}

type NodeQueue struct {
	nodes []*Tree
	cur int
}

func (nq *NodeQueue) Pop() *Tree {
	node := nq.nodes[nq.cur]
	nq.cur++
	return node
}

func (nq *NodeQueue) Push(node *Tree) {
	nq.nodes = append(nq.nodes, node)
}

func NewNodeQueue() *NodeQueue{
	nq := &NodeQueue{
		nodes: []*Tree{},
		cur : 0,
	}
	return nq
}

func PreOrder(node *Tree) {
	if node == nil {
		return
	}
	fmt.Print(node.val + " ")
	PreOrder(node.left)
	PreOrder(node.right)
}

func InOrder(node *Tree) {
	if node == nil {
		return
	}
	InOrder(node.left)
	fmt.Print(node.val + " ")
	InOrder(node.right)
}

func PostOrder(node *Tree) {
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	fmt.Print(node.val + " ")
}
// 按层的顺序遍历树节点
func SequentialOrder(node *Tree) {
	nq := NewNodeQueue()
	nq.Push(node)
	for nq.cur < len(nq.nodes) {
		n := nq.Pop()
		fmt.Print(n.val + " ")
		if n.left != nil {
			nq.Push(n.left)
		}
		if n.right != nil {
			nq.Push(n.right)
		}
	}
}
// 求给定二叉树的高度
func getHeightOfBinaryTree(root *Tree) float64 {
	return math.Max(float64(getHeightOfBinaryLeftSubTree(root.left, 0)),
					float64(getHeightOfBinaryRightSubTree(root.right, 0))) + float64(1)
}

func getHeightOfBinaryLeftSubTree(node *Tree, height int) int {
	if node == nil {
		return height
	}
	height++
	if node.left == nil {
		return height
	} else {
		return getHeightOfBinaryLeftSubTree(node.left, height)
	}
}

func getHeightOfBinaryRightSubTree(node *Tree, height int) int {
	if node == nil {
		return height
	}
	height++
	if node.right == nil {
		return height
	} else {
		return getHeightOfBinaryRightSubTree(node.right, height)
	}
}


