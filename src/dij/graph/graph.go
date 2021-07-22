package graph

type Edge struct {
	S int
	T int
	W int
}


type Graph struct {
	Adj [][]*Edge
	V int
}

func InitialGraph(v int) *Graph {
	g := new(Graph)
	g.Adj = make([][]*Edge, v)
	g.V = v
	return g
}

func (g *Graph) AddEdge(s, t, w int) {
	g.Adj[s] = append(g.Adj[s], &Edge{s, t, w})
}
