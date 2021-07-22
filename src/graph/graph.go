package graph

import "fmt"

type Graph struct {
	v int		// 顶点数
	adj [][]int	//用于存储图的邻接表
}

func initialGraph(v int) *Graph {
	g := &Graph{
		v: v,
		adj: make([][]int, v),
	}
	return g
}

func (g *Graph) addEdge(s, t int) {
	g.adj[s] = append(g.adj[s], t)
	g.adj[t] = append(g.adj[t], s)
}

type Queue struct {
	items []int
	capacity int
	p int
}

func (q *Queue) put(point int) {
	if q.capacity == len(q.items) {
		return
	}
	q.items = append(q.items, point)
}

func (q *Queue) pop() int {
	if len(q.items) == 0 || q.p == len(q.items) {
		return -1
	}
	point := q.items[q.p]
	q.p++
	return point
}

// Breadth-First-Search
func (g *Graph) bfs(s, t int) int{
	if s == t {
		return 0
	}
	visited := make([]bool, g.v)
	q := Queue{
		capacity: g.v,
		p: 0,
	}
	q.put(s)
	visited[s] = true
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}
	for q.p < len(q.items) {
		cur := q.pop()
		for i := 0; i < len(g.adj[cur]); i++ {
			w := g.adj[cur][i]
			if !visited[w] {
				prev[w] = cur
				if w == t {
					return print(prev, s, t, 0) - 1
				}
				q.put(w)
				visited[w] = true
			}
		}
	}
	return -1
}

var found = false
func (g *Graph) recurDFS(s, t int, prev []int, visited []bool) {
	if found {
		return
	}
	if s == t {
		found = true
		return
	}
	visited[s] = true
	for i := 0; i < len(g.adj[s]); i++ {
		if found {
			return
		}
		q := g.adj[s][i]
		if !visited[q] {
			prev[q] = s
			g.recurDFS(q, t, prev, visited)
		}
	}
}

func (g *Graph) NonRecurDFS(s, t int, prev []int, visited []bool) {
	if found {
		return
	}
	if s == t {
		found = true
		return
	}
	visited[s] = true
	for true {
		allVisited := true
		k := 0
		for true {
			if k < len(g.adj[s]) {
				q := g.adj[s][k]
				if !visited[q] {
					prev[q] = s
					if q == t {
						found = true
						return
					}
					visited[q] = true
					s = q
					break
				}
				k++
				continue
			}
			s = prev[s]
			break
		}
		for i := 0; i < g.v; i++ {
			if !visited[i] {
				allVisited = false
				break
			}
		}
		if allVisited {
			fmt.Println("invaid parameters!")
			return
		}
	}
}

func print(prev []int, s, t, v int) int{
	v++
	if prev[t] != -1 && t != s {
		v = print(prev, s, prev[t], v)
	}
	//fmt.Printf("%v ", t)
	return v
}

func printRoute(prev []int, s, t int) {
	if prev[t] != -1 && s != t {
		printRoute(prev, s, prev[t])
	}
	fmt.Printf("%v ", t)
}



