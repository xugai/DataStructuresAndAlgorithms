package bm

import (
	"fmt"
	"testing"
)

func TestBm(t *testing.T) {
	a := "acdef"
	b := "acd"
	fmt.Println(bm(a, b))
}
