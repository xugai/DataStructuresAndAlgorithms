package arrayqueue

import "fmt"

type Queue struct {
	items []int
	capacity int
	head int
	tail int
}

func initQueue(capacity int) *Queue {
	q := Queue{
		items: make([]int, capacity),
		capacity: capacity,
	}
	return &q
}

func (q *Queue) inqueue(elem int) bool {
	if q.tail >= q.capacity {
		// 若队列已满，但有出队操作，则将当前剩余的队列中的元素搬移至队列头
		if q.head > 0 && q.head < q.tail {
			for i := q.head; i < q.tail; i++ {
				q.items[i - q.head] = q.items[i]
				q.items[i] = 0
			}
			q.tail = q.capacity - q.head
			q.head = 0
			q.items[q.tail] = elem
			q.tail++
			return true
		} else {
			fmt.Println("current queue is full!")
			return false
		}
	}
	q.items[q.tail] = elem
	q.tail++
	return true
}

func (q *Queue) dequeue() int {
	if q.head == q.tail {
		return -1
	}
	elem := q.items[q.head]
	q.items[q.head] = 0
	q.head++
	return elem
}

func (q *Queue) print() {
	if q.head == q.tail {
		fmt.Println("current queue is empty!")
		return
	}
	for i := 0; i < len(q.items); i++ {
		fmt.Printf("%v ", q.items[i])
	}
	fmt.Println()
}
