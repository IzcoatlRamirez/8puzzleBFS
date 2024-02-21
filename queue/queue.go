package queue

import (
	"container/list"
	"github.com/IzcoatlRam/puzzle8/node"
)

type Queue struct{
	q *list.List
}

func NewQueue() *Queue {
	return &Queue{
		q : list.New(),
	}
}

func (q *Queue) Encolar(n *node.Node) {
	q.q.PushBack(n)
}

func (q *Queue) Desencolar() *node.Node {
	e := q.q.Front()
	q.q.Remove(e)
	return e.Value.(*node.Node)
}

func (q *Queue) Len() int {
	return q.q.Len()
}