package binarysearchtree

import "fmt"

// 二叉查找树
type Node struct {
	data int
	left *Node
	right *Node
}
// 查找二叉查找树中的某个节点
func SearchNode(root *Node, target *Node) *Node {
	if target == nil || root == nil {
		return nil
	}
	if root.data > target.data {
		return SearchNode(root.left, target)
	} else if root.data < target.data {
		return SearchNode(root.right, target)
	} else {
		return root
	}
}
// 往二叉查找树中插入一个节点
func InsertNode(root *Node, node *Node) bool {
	if root == nil || node == nil || root.data == node.data {
		return false
	}
	if root.data > node.data {
		if root.left == nil {
			root.left = node
			return true
		} else {
			return InsertNode(root.left, node)
		}
	} else if root.data < node.data {
		if root.right == nil {
			root.right = node
			return true
		} else {
			return InsertNode(root.right, node)
		}
	}
	return false
}
// 往二叉查找树中删除一个结点
/*
	如果被删除的节点只有其自己本身而已；
	如果被删除的节点只有一个子节点（左节点或者右节点）;
	如果被删除的节点都有左右子节点；
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			root = root.Right
		}  else if root.Right == nil {
			root = root.Left
		} else {
			root.Val = findSuccessor(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findSuccessor(node *TreeNode) int {
	if node == nil {
		return -1
	}
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

// 往二叉搜索树中插入一个节点
func insertNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		root = &TreeNode{key, nil,nil}
		return root
	}
	if root.Val > key {
		root.Left = insertNode(root.Left, key)
	} else if root.Val < key {
		root.Right = insertNode(root.Right, key)
	}
	return root
}

func traverseTree(root *TreeNode) {
	treeNodeCollector := []*TreeNode{}
	treeNodeCollector = append(treeNodeCollector, root)
	var node *TreeNode
	p := 1
	k := 0
	for p > 0 {
		node = treeNodeCollector[k]
		k++
		p--
		if node != nil {
			fmt.Printf("%v ", node.Val)
		} else {
			fmt.Print(nil)
			fmt.Print(" ")
			continue
		}
		if node.Left == nil && node.Right == nil {

		} else {
			treeNodeCollector = append(treeNodeCollector, node.Left, node.Right)
			p += 2
		}
	}
}