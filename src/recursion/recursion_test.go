package recursion

import "testing"

/*
	递归适用的场景:
	1/ 一个问题是否能拆分为多个具有同样解法的子问题
	2/ 这些拆分出来的子问题是否具有终止条件
*/

func fn(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fn(n - 1) + fn(n - 2)
}


func TestPermutations(t *testing.T) {
	arr := []int{1, 2, 4}
	permutations(arr, len(arr))
}
