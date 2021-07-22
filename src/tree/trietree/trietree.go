package trietree

// Trie Tree, 其经典的使用场景是，根据搜索关键词进行提示，也就是根据前缀子串，查询出多个与前缀子串匹配的字符串
//type TrieTreeNode struct {
//	data string
//	children []*TrieTreeNode
//	isEndingChar bool
//}
//
//func NewTrieTree(data string) *TrieTreeNode {
//	root := new(TrieTreeNode)
//	root.data = data
//	root.children = make([]*TrieTreeNode, 26)
//	return root
//}
//
//var root = NewTrieTree("/")
//func insertTrieTree(str string) {
//	p := root
//	for i := 0; i < len(str); i++ {
//		index := int(rune(str[i])) - int('a')
//		if p.children[index] == nil {
//			trieTreeNode := NewTrieTree(string(str[i]))
//			p.children[index] = trieTreeNode
//		}
//		p = p.children[index]
//	}
//	p.isEndingChar = true
//}
//
//func searchTrieTree(str string) bool {
//	if len(str) == 0 {
//		return false
//	}
//	p := root
//	for i := 0; i < len(str); i++ {
//		index := int(rune(str[i])) - int('a')
//		if p.children[index] == nil {
//			return false
//		}
//		p = p.children[index]
//	}
//	if !p.isEndingChar {
//		return false
//	} else {
//		return true
//	}
//}
