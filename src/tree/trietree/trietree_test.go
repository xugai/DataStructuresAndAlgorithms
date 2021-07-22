package trietree

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	words := []string{
		"time",
		"me",
		"bell",
	}
	sortWords(words)
	fmt.Println(insertTrieTreeNode(words))
}

func sortWords(words []string) {
	tmp := ""
	for i := 0; i < len(words) - 1; i++ {
		if len(words[i]) < len(words[i + 1]) {
			tmp = words[i]
			words[i] = words[i + 1]
			words[i + 1] = tmp
		}
	}
}
