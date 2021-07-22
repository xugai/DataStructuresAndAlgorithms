package lcs

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	a := "mtacnu"
	b := "mitcmu"
	fmt.Println(lcs(a, b, len(a), len(b)))
}

func TestMaxSeq(t *testing.T) {
	numStr := "2936517"
	maxNumSeq := ""
	for i := 0; i < len(numStr); i++ {
		maxNumSeq = string(numStr[i])
		longestIncreaseSubString(i, numStr, maxNumSeq)
	}
	fmt.Println(maxSequence)
}