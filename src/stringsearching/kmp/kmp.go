package kmp

// a、b分别表示主串与模式串，n、m表示主串与模式串的长度
func kmp(a, b string, n, m int) int {
	if len(a) == 0 || len(b) == 0 {
		return -1
	}
	j := 0
	next := generateNext(b, m)
	for i := 0; i < n; i++ {
		for j > 0 && a[i] != b[j] {
			j = next[j - 1] + 1
		}
		if a[i] == b[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func generateNext(b string, m int) []int {
	if len(b) == 0 {
		return nil
	}
	next := make([]int, m)
	next[0] = -1
	k := -1
	for i := 1; i < m; i++ {
		for k != -1 && b[i] != b[k + 1] {
			k = next[k]
		}
		if b[i] == b[k + 1] {
			k++
		}
		next[i] = k
	}
	return next
}
