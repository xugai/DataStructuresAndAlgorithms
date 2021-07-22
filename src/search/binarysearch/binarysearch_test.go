package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(BinarySearchInLoopSeqArr(arr, 8))
}
