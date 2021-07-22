package toposorting

var adj []*linkedList

type linkedList struct {
	items []int
}

type queue struct {
	items []int
	p int
}

func (l *linkedList) add(val int) {
	l.items = append(l.items, val)
}

func initGraph(v int) {
	adj = make([]*linkedList, v + 1)
	for i := 1; i <= v; i++ {
		adj[i] = new(linkedList)
	}
}

func addEdge(s, t int) {
	adj[s].add(t)
}
