package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 8, 6, 1, 7}
	BubblesortWithNoFlag(&arr)
	fmt.Println(arr)
}

func BenchmarkBubbleSortWithFlag(b *testing.B) {
	b.ResetTimer()
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Bubblesort(&arr)
	}
	b.StopTimer()
}

func BenchmarkBubbleSortWithNoFlag(b *testing.B) {
	b.ResetTimer()
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		BubblesortWithNoFlag(&arr)
	}
	b.StopTimer()
}
