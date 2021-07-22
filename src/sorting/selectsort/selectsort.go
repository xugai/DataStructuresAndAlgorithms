package selectsort

func SelectSorting(arr *[]int) {
	arrLength := len(*arr)
	if arrLength <= 1 {
		return
	}
	for i := 0; i < arrLength; i++ {
		head := i
		min := (*arr)[i]
		tmp := 0
		tmpIndex := i
		for j := head; j < arrLength; j++ {
			if (*arr)[j] < min {
				min = (*arr)[j]
				tmpIndex = j
			}
		}
		tmp = (*arr)[i]
		(*arr)[i] = min
		(*arr)[tmpIndex] = tmp
	}
}
