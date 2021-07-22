package mergesort

import (
	"fmt"
	"testing"
)

func TestMergesort(t *testing.T) {
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	result := MergeSort(arr)
	fmt.Println(result)
}

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	subArr := arr[0:3]
	fmt.Println(subArr)
	fmt.Println(subArr[1:])
	//fmt.Println(len(subArr))
}
