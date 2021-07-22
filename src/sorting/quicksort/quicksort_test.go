package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 7, 3, 8, 9, 2}
	QuickSort(&arr)
	fmt.Println(arr)
}

func TestFindNumber(t *testing.T) {
	arr := []int{4, 2, 5, 12, 3}
	result := FindNumber(arr, 1)
	fmt.Println(result)
}
