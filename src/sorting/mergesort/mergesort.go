package mergesort

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	q := len(arr) / 2
	left := MergeSort(arr[0:q])
	right := MergeSort(arr[q:])
	return MergeSortFinal(left, right)
}

func MergeSortFinal(left, right []int) []int {
	i, j:= 0, 0
	l, r := len(left), len(right)
	tmp := make([]int, 0)
	for i < l && j < r {
		if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
			continue
		}
		tmp = append(tmp, right[j])
		j++
	}
	for i < l {
		tmp = append(tmp, left[i])
		i++
	}
	for j < r {
		tmp = append(tmp, right[j])
		j++
	}
	fmt.Println(tmp)
	return tmp
}
