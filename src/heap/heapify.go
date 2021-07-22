package heap

import "math"

type Heap struct {
	content []int	// 用于存储数据的数组
	capacity int	// 堆的容量
	count int		// 当前堆中的数组存储了几个元素
}

func InitialHeap(capacity int) *Heap{
	return &Heap{
		content: make([]int, capacity + 1),
		capacity: capacity,
		count: 0,
	}
}

func insertNode(heap *Heap, data int) {
	if heap.capacity == 0 || heap.count == heap.capacity {
		return
	}
	heap.count++
	heap.content[heap.count] = data
	tmp := 0
	i := heap.count
	for i / 2 > 0 && heap.content[i] > heap.content[i / 2] {
		tmp = heap.content[i / 2]
		heap.content[i / 2] = heap.content[i]
		heap.content[i] = tmp
		i = i / 2
	}
}

func deleteTopOfNode(heap *Heap, data int) {
	if len(heap.content) == 0 || heap.count == 0 {
		return
	}
	heap.content[1] = data
	heap.count--
	i := 1
	maxPosition := 0
	for true {
		maxPosition = i
		if 2 * i <= len(heap.content) && heap.content[2 * i] > heap.content[i] {
			maxPosition = 2 * i
		}
		if 2 * i + 1 <= len(heap.content) && heap.content[2 * i + 1] > heap.content[maxPosition] {
			maxPosition = 2 * i + 1
		}
		if maxPosition == i {
			break
		}
		// todo swap(heap.content, i, maxPosition)
		i = maxPosition
	}
}

// 建堆
func buildHeap(points []*Point, n int) {
	if n <= 0 {
		return
	}
	var tmp *Point
	maxPosition := -1
	for i := n / 2; i >= 1; i-- {
		for true {
			maxPosition = i
			if 2 * i <= n && points[2 * i].val > points[i].val {
				maxPosition = 2 * i
			}
			if 2 * i + 1 <= n && points[2 * i + 1].val > points[maxPosition].val {
				maxPosition = 2*i + 1
			}
			if maxPosition == i {
				break
			}
			tmp = points[i]
			points[i] = points[maxPosition]
			points[maxPosition] = tmp
			i = maxPosition
		}
	}
}

// 堆化
func heapify(points []*Point, n int) {
	if n <= 0 {
		return
	}
	var tmp *Point
	maxPosition, i := 1, 1
	for true {
		maxPosition = i
		if 2 * i <= n && points[2 * i].val > points[i].val {
			maxPosition = 2 * i
		}
		if 2 * i + 1 <= n && points[2 * i + 1].val > points[maxPosition].val {
			maxPosition = 2 * i + 1
		}
		if maxPosition == i {
			break
		}
		tmp = points[i]
		points[i] = points[maxPosition]
		points[maxPosition] = tmp
		i = maxPosition
	}
}

// 堆排序
func heapSorting(arr []int, n int) {
	if n == 0 {
		return
	}
	for n > 1 {
		tmp := arr[n]
		arr[n] = arr[1]
		arr[1] = tmp
		n--
		//heapify(arr, n)
	}
}

type Point struct {
	val float64
	index int
}


func kClosest(points [][]int, k int) [][]int {
	sqrtRecord := make([]float64, len(points))
	topKHeap := make([]*Point, k + 1)
	result := [][]int{}
	for i := 0; i < len(points); i++ {
		sqrtRecord[i] = math.Sqrt(float64(points[i][0] * points[i][0]) + float64(points[i][1] * points[i][1]))
		if i < k {
			topKHeap[i+1] = &Point{
				val:   sqrtRecord[i],
				index: i,
			}
		}
	}
	buildHeap(topKHeap, k)
	for i := k; i < len(sqrtRecord); i++ {
		if topKHeap[1].val > sqrtRecord[i] {
			topKHeap[1].val = sqrtRecord[i]
			topKHeap[1].index = i
			heapify(topKHeap, k)
		}
	}
	for i, p := range topKHeap {
		if i == 0 {
			continue
		}
		result = append(result, points[p.index])
	}
	return result
}

