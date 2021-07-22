package recursion

import "fmt"

func permutations(arr []int, k int) {
	if len(arr) == 0 {
		return
	}

	if k == 1 {
		for i := len(arr) - 1; i >= 0; i-- {
			fmt.Printf("%v ", arr[i])
		}
		fmt.Println()
		return
	}

	tmp := 0
	for i := 0; i < k; i++ {
		tmp = arr[i]
		arr[i] = arr[k - 1]
		arr[k - 1] = tmp

		permutations(arr, k - 1)

		tmp = arr[i]
		arr[i] = arr[k - 1]
		arr[k - 1] = tmp
	}
}
