package graph

import "sync"

type Stack struct {
	nodes []Node
	mutex sync.RWMutex
}

func (s *Stack)New() *Stack{
	s.mutex.Lock()
	s.nodes = []Node{}
	s.mutex.Unlock()

	return s
}

func (s *Stack)Add(n Node){
	s.mutex.Lock()
	s.nodes = append(s.nodes, n)
	s.mutex.Unlock()
}

func (s *Stack)Pop() *Node{
	s.mutex.Lock()
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0:len(s.nodes)-1]
	s.mutex.Unlock()
	return &node
}

func (s *Stack)IsEmpty() bool{
	s.mutex.RLock()
    defer s.mutex.RUnlock()
	return (len(s.nodes) == 0)
}