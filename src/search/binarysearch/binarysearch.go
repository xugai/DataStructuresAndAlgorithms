package binarysearch

// 最经典的二分查找——查找等于给定值的元素
func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 查找第一个等于给定值的元素（数组里面存在重复元素）
func BinarySearchForFirstEqualElem(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid - 1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 查找最后一个等于给定值的元素（数组里面存在重复元素）
func BinarySearchForLastEqualElem(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == len(arr) - 1 || arr[mid + 1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素（数组里面存在重复元素）
func BinarySearchForFirstGreaterOrEqualElem(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid - 1] < target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素（数组里面存在重复元素）
func BinarySearchForLastLessOrEqualElem(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else {
			if mid == len(arr) - 1 || arr[mid + 1] > target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 在一个有序循环数组里面，查找等于给定值的元素，例如[4,5,6,1,2,3]
func BinarySearchInLoopSeqArr(arr []int, target int) int {
	// 默认从小到大排列
	p := 0
	for i := 1; i < len(arr) - 1; i++ {
		if arr[i] > arr[i - 1] {
			continue
		}
		p = i // 找到分区点,[0:i-1],[i:]
		break
	}
	if result := BinarySearch(arr[0:p], target); result != -1 {
		return result
	}
	if result := BinarySearch(arr[p:], target); result != -1 {
		return p + result
	}
	return -1
}






