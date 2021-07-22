package trietree

type TrieTreeNode struct {
	val string
	children []*TrieTreeNode
}

func NewTrieTreeNode() *TrieTreeNode {
	node := new(TrieTreeNode)
	node.children = make([]*TrieTreeNode, 26)
	return node
}

func insertTrieTreeNode(words []string) int {
	root := NewTrieTreeNode()
	length := 0
	for i := 0; i < len(words); i++ {
		p := root
		isNewWord := false
		word := words[i]
		for j := len(word) - 1; j >= 0; j-- {
			index := int(rune(word[j])) - int('a')
			if p.children[index] == nil {
				isNewWord = true
				p.children[index] = NewTrieTreeNode()
			}
			p = p.children[index]
		}
		if isNewWord {
			length += len(word) + 1
		}
	}
	return length
}
