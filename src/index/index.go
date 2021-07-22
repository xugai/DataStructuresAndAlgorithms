package index

/*
	数据库B+树索引，叶子节点存放数据，非叶子节点存放索引，用来指向下一个节点
	叶子节点通过链表连接起来
	B+树索引，节点内容存放在磁盘里，通过尽可能地降低树的高度，来达到减少磁盘IO次数，提高读写效率
*/

// B+树非叶子节点
type BPlusTreeNode struct {
	m int
	keywords []int
	children []*BPlusTreeNode
}

// B+树叶子节点
type BPlusTreeLeafNode struct {
	k int
	keywords []int
	dataAddress []int64
	prev *BPlusTreeLeafNode		// 前驱
	next *BPlusTreeLeafNode		// 后继
}


