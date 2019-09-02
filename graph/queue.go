package graph

import "sync"

type Queue struct {
	nodes []Node
	mutex sync.RWMutex
}

func (q *Queue)New() *Queue{
	q.mutex.Lock()
	q.nodes = []Node{}
	q.mutex.Unlock()

	return q
}

func (q *Queue)Enqueue(n Node){
	q.mutex.Lock()
	q.nodes = append(q.nodes, n)
	q.mutex.Unlock()
}

func (q *Queue)Dequeue()*Node{
	q.mutex.Lock()
	node := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.mutex.Unlock()
	return &node
}

func (q *Queue)IsEmpty() bool{
	q.mutex.RLock()
    defer q.mutex.RUnlock()
	return (len(q.nodes) == 0)
}