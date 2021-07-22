package quicksort

import "fmt"

func QuickSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	q := Partition(arr)
	left := (*arr)[0:q]
	QuickSort(&left)
	right := (*arr)[q:]
	QuickSort(&right)
}

func Partition(arr *[]int) int {
	pivot := (*arr)[len(*arr) - 1]
	i := 0
	tmp := 0
	for j := 0; j < len(*arr) - 1; j++ {
		if (*arr)[j] < pivot {
			tmp = (*arr)[j]
			(*arr)[j] = (*arr)[i]
			(*arr)[i] = tmp
			i++
		}
	}
	tmp = (*arr)[i]
	(*arr)[i] = pivot
	(*arr)[len(*arr) - 1] = tmp
	return i
}

func FindNumber(arr []int, k int) int {
	if len(arr) < k {
		fmt.Println("invalid params")
		return -1
	}
	q := PartitionForFindNumber(arr)
	if q + 1 == k {
		return arr[q]
	} else if k > q + 1 {
		return FindNumber(arr[q:], k)
	} else{
		return FindNumber(arr[0:q], k)
	}
}

func PartitionForFindNumber(arr []int) int {
	pivot := arr[len(arr) - 1]
	i := 0
	tmp := 0
	for j := 0; j < len(arr) - 1; j++ {
		if arr[j] > pivot {
			tmp = arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			i++
		}
	}
	tmp = arr[i]
	arr[i] = pivot
	arr[len(arr) - 1] = tmp
	return i
}
