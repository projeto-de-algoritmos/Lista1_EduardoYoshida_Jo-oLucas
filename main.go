package main

import (
    "fmt"
    "grafos/graph"
    "bufio"
    "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    var option string
	var g graph.Graph

    for{
        fmt.Println("\n ------------------")
        fmt.Println("Selecione uma opção: ")
        fmt.Print(" 1 - Inserir vértice no Grafo \n 2 - Inserir aresta no grafo \n 3 - Utilizar grafo pré-definido \n 4 - Imprimir lista de adjacência \n 5 - Buscar nó(BFS) \n")
        scanner.Scan()
        option = scanner.Text()
        switch option{
        case "1":
            fmt.Print("Insira o nome do nó: ")
            scanner.Scan()
            name := scanner.Text()
            node := graph.Node{Key: name}
            g.AddNode(&node)
        case "2":
            fmt.Print("Insira o primeiro nó: ")
            scanner.Scan()
            n1 := g.GetNode(scanner.Text())
            fmt.Print("Insira o segundo nó: ")
            scanner.Scan()
            n2 := g.GetNode(scanner.Text())
            g.AddEdge(n1, n2)
        case "3":
            fillGraph(&g)
            fmt.Println("Grafo carregado com sucesso")
        case "4":
            g.String()
        case "5":
            fmt.Print("Insira o nó pelo qual deseja iniciar a busca: ")
            scanner.Scan()
            n1 := g.GetNode(scanner.Text())
            fmt.Print("Insira o nó desejado: ")
            scanner.Scan()
            n2 := g.GetNode(scanner.Text())
            if(n1 == nil || n2 == nil){
                fmt.Println("Nós inválidos")
            }else{
                fmt.Println("\n ------------------\n BFS Traversal Order")
                prev := g.BFS(n1, func(n *graph.Node){
                        fmt.Printf("%v\n", n)
                })
                path := BFSUtil(n1, n2, prev, &g)
                fmt.Println("\n ------------------")
                fmt.Printf("O caminho mais curto entre %v e %v é: \n", n1, n2)
                printPath(path)
            }
        }
    }
}

func printPath(path []*graph.Node){
    str := path[len(path) -1].Key
    for i:= len(path) -2; i>=0; i--{
        str += "->" + path[i].Key
    }
    fmt.Println(str)
}

func fillGraph(g *graph.Graph){
	nA := graph.Node{Key: "Joao", Value: 1}
    nB := graph.Node{Key: "Eduardo", Value: 2}
    nC := graph.Node{Key: "Paulo", Value: 3}
    nD := graph.Node{Key: "Adriane", Value: 4}
    nE := graph.Node{Key: "Maurício", Value: 5}
    nF := graph.Node{Key: "Milene", Value: 6}
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

func BFSUtil(n1, n2 *graph.Node, prev map[*graph.Node]*graph.Node, g *graph.Graph) []*graph.Node{
    current := n2
    path := make([]*graph.Node, 0)
    path = append(path, n2)
    for prev[current] != nil{
        path = append(path, prev[current])
        // current = prev[current] TODO FIX FOR SOME REASON CURRENT IS POINTING TO A NULL POINTER
        current = g.GetNode(prev[current].Key) 
    }
    return path
}
