package queue

import "priority-queue/list"

type (
	PQueue interface {
		Push(x int, priority int)
		Pop() int
		Shift() int
		GetMax() int
		GetMin() int
	}

	pqueue struct {
		list list.List
	}
)

func NewPQueue() PQueue {
	return &pqueue{list: list.NewList()}
}

func (q *pqueue) Push(x int, priority int) {
	q.list.Add(x, priority)
}

func (q *pqueue) Pop() int {
	m := q.list.GetMin()
	q.list.Delete(m)

	return m.Val
}

func (q *pqueue) Shift() int {
	m := q.list.GetMax()
	q.list.Delete(m)

	return m.Val
}

func (q *pqueue) GetMax() int {
	return q.list.GetMax().Val
}

func (q *pqueue) GetMin() int {
	return q.list.GetMin().Val
}
