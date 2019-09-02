package main

import (
    "fmt"
    "grafos/graph"
)

func main(){
	var g graph.Graph
    fillGraph(&g)
    
	g.BFS(func(n *graph.Node){
		fmt.Printf("%v\n", n)
	})
}

func fillGraph(g *graph.Graph){
	nA := graph.Node{Key: "A", Value: 1}
    nB := graph.Node{Key: "B", Value: 2}
    nC := graph.Node{Key: "C", Value: 3}
    nD := graph.Node{Key: "D", Value: 4}
    nE := graph.Node{Key: "E", Value: 5}
    nF := graph.Node{Key: "F", Value: 6}
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)
}