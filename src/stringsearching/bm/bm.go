package bm

import "math"

/*
	字符串匹配算法——BM算法
	- 坏字符规则
		- 模式串与主串进行匹配，从模式串的最后一位开始匹配。当发现第一位不匹配的字符时，该不匹配的字符就是坏字符，记录该坏字符在模式串中的下标为x。
		- 然后拿着这个坏字符，在模式串中查找是否有其他字符跟坏字符一样。
		- 如果找得到，则记录当前找到的字符的下标位置，设为y。然后x-y就是模式串需要滑动的位数
		- 如果找不到，则将y设为-1。
	- 好后缀规则
		- 同样从模式串的最后一位开始进行匹配，如果查找过程中有匹配的字符，然后第一次找到坏字符，那么记录已经查找到的匹配的字符串，这个匹配的字符串即为好后缀。
		- 然后拿着找到的好后缀，在模式串中查找是否有其他与好后缀一样的字符串，如果有，则将这个匹配的字符串的起始位置标记为r，那么模式串下一次需要滑动m-r+1个位置
		- 如果模式串中没有与好后缀一样的字符串，那就寻找好后缀中能够与模式串的前缀子串匹配的最长后缀子串，然后记录这最长后缀子串的下标位置r
		- 那么模式串下一次就需要往后滑动r个位置了。
		- 如果模式串中找不到任何一个与好后缀中的后缀子串匹配的前缀子串，那么就返回m，说明模式串下一次往后滑动m位.
*/

// a是主串，b是模式串
const SIZE = 256
func generateBC(b string, bc []int) {
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}
	for i := 0; i < len(b); i++ {
		bc[int(rune(b[i]))] = i
	}
}

func bm(a, b string) int  {
	if len(a) == 0 || len(b) == 0 {
		return -1
	}
	bc := make([]int, SIZE)
	generateBC(b, bc)
	suffix := make([]int, len(b))
	prefix := make([]bool, len(b))
	generateGS(b, suffix, prefix)
	z := 0
	for z <= len(a) - len(b) {
		j := 0
		for j = len(b) - 1; j >= 0; j-- {
			if b[j] != a[j + z] {
				break
			}
		}
		if j < 0 {
			return z
		}
		x := j - bc[int(rune(a[j + z]))]
		y := 0
		if j < len(b) - 1 {
			// 说明出现好后缀
			y = moveByGS(j, b, suffix, prefix)
		}
		z = z + int(math.Max(float64(x), float64(y)))
	}
	return -1
}

func generateGS(b string, suffix []int, prefix []bool) {
	if len(b) == 0 {
		return
	}
	for i := 0; i < len(b); i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < len(b) - 1; i++ {
		j := i
		k := 0
		for j >= 0 && b[j] == b[len(b) - 1 - k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j < 0 {
			prefix[k] = true
		}
	}
}

func moveByGS(j int, b string, suffix []int, prefix []bool) int {
	k := len(b) - 1 - j  // 求出好后缀的长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r < len(b); r++ {
		if prefix[len(b) - r] {
			return r
		}
	}
	return len(b)
}


