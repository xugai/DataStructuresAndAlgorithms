package insertsort

// 交换次数 = 逆序度 = 满有序度 - 有序度
// 满有序度 = n * (n - 1) / 2
func InsertSorting(arr *[]int) {
	arrLength := len(*arr)
	if arrLength <= 1 {
		return
	}
	for i := 1; i < arrLength; i++ {
		value := (*arr)[i]
		j := i - 1
		for ; j >= 0; j-- {
			if (*arr)[j] > value {
				(*arr)[j + 1] = (*arr)[j]
			} else {
				break
			}
		}
		(*arr)[j + 1] = value
	}
}
