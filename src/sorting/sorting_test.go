package sorting

import (
	"DataStructuresAndAlgorithms/src/sorting/bubblesort"
	"DataStructuresAndAlgorithms/src/sorting/insertsort"
	"DataStructuresAndAlgorithms/src/sorting/selectsort"
	"fmt"

	//"fmt"
	"testing"
)

func BenchmarkWithBubbleSorting(b *testing.B) {
	b.ResetTimer()
	arr := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		bubblesort.Bubblesort(&arr)
	}
	//fmt.Println(arr)
	b.StopTimer()
}

func BenchmarkWithInsertSorting(b *testing.B) {
	b.ResetTimer()
	arr := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		insertsort.InsertSorting(&arr)
	}
	//fmt.Println(arr)
	b.StopTimer()
}

func BenchmarkWithSelectSorting(b *testing.B) {
	b.ResetTimer()
	arr := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		selectsort.SelectSorting(&arr)
	}
	//fmt.Println(arr)
	b.StopTimer()
}

func TestDevide(t *testing.T) {
	q, r := 0, 8
	fmt.Println((q + r) / 2)
}

/*
	对于像slice[low:high:max]这样的切片声明，先看low跟max组合成的切片是什么样子，然后再看low跟high组合成的切片
*/
func TestSliceView(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]  // newslice len is 3-1=2, cap is 5-1=4
	fmt.Println(newSlice)
	newSlice = append(newSlice, 60)
	fmt.Printf("newslice: %v\n", newSlice)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("slice[2:3:4]: %v\n", slice[2:3:4])
	newSlice = append(newSlice, []int{70, 80}...)
	fmt.Println("after appending...")
	fmt.Printf("newslice: %v\n", newSlice)
	fmt.Printf("slice: %v\n", slice)
}
