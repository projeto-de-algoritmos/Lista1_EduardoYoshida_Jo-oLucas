package graph

import(
    "fmt"
    "sync"
)
type Node struct{
    Key string
    Value int
}

func (n *Node) String() string {
    return fmt.Sprintf("%v", n.Key)
}

type Graph struct{
	nodes []*Node
    edges map[Node][]*Node
    mutex  sync.RWMutex
}

func (g *Graph) AddNode(n *Node){
    g.mutex.Lock()
    g.nodes = append(g.nodes, n)
    g.mutex.Unlock()
}

func (g *Graph) AddEdge(node1, node2 *Node){
    g.mutex.Lock()
	if g.edges == nil{
		g.edges = make(map[Node][]*Node)
	}
	// non directional graph implementation, so edges point both ways
	g.edges[*node1] = append(g.edges[*node1], node2)
    g.edges[*node2] = append(g.edges[*node2], node1)
    g.mutex.Unlock()
}

func (g *Graph) String() {
    g.mutex.RLock()
    s := ""
    for i := 0; i < len(g.nodes); i++ {
        s += g.nodes[i].String() + " -> "
        near := g.edges[*g.nodes[i]]
        for j := 0; j < len(near); j++ {
            s += near[j].String() + " "
        }
        s += "\n"
    }
    fmt.Println(s)
    g.mutex.RUnlock()
}

func (g *Graph) BFS(f func(*Node)) {
    g.mutex.RLock()
    q := Queue{}
    q.New()
    n := g.nodes[0]
    q.Enqueue(*n)
    visited := make(map[*Node]bool)
    for {
        if q.IsEmpty() {
            break
        }
        node := q.Dequeue()
        visited[node] = true
        near := g.edges[*node]

        for i := 0; i < len(near); i++ {
            j := near[i]
            if !visited[j] {
                q.Enqueue(*j)
                visited[j] = true
            }
        }
        if f != nil {
            f(node)
        }
    }
    g.mutex.RUnlock()
}

