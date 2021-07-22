package kmp

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	a := "abacba"
	b := "acb"
	fmt.Println(kmp(a, b, 6, 3))
}
