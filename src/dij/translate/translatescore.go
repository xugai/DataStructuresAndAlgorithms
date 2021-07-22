package translate

type Queue struct {
	scores []int
}

func initialQueue() *Queue{
	q := new(Queue)
	q.scores = []int{}
	q.scores = append(q.scores, -1)
	return q
}

func (q *Queue) Put(score int) {
	q.scores = append(q.scores, score)
	n := len(q.scores) - 1
	for n / 2 > 0 && q.scores[n] > q.scores[n / 2] {
		tmp := q.scores[n]
		q.scores[n] = q.scores[n / 2]
		q.scores[n / 2] = tmp
		n = n / 2
	}
}

func (q *Queue) Pop() int {
	score := q.scores[1]
	q.scores[1] = q.scores[len(q.scores) - 1]
	maxPosition, i := 1, 1
	n := len(q.scores)
	for true {
		maxPosition = i
		if 2 * i < n && q.scores[2 * i] > q.scores[i] {
			maxPosition = 2 * i
		}
		if 2 * i + 1 < n && q.scores[2 * i + 1] > q.scores[maxPosition] {
			maxPosition = 2 * i + 1
		}
		if maxPosition == i {
			break
		}
		tmp := q.scores[maxPosition]
		q.scores[maxPosition] = q.scores[i]
		q.scores[i] = tmp
		i = maxPosition
	}
	return score
}

var scores = initialQueue()
// calculateScore(0, mem, 0)
func calculateScore(n int, mem [][]int, score int) {
	if n == 3 {
		scores.Put(score)
		return
	}
	for i := 0; i < len(mem[n]); i++ {
		calculateScore(n + 1, mem, score + mem[n][i])
	}
}
