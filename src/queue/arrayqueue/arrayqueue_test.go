package arrayqueue

import "testing"

func TestArrayQueue(t *testing.T) {
	q := initQueue(5)
	q.dequeue()
	q.inqueue(1)
	q.inqueue(2)
	q.dequeue()
	q.print()
}
