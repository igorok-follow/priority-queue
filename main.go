package main

import (
	"log"
	"priority-queue/queue"
)

func main() {
	q := queue.NewPQueue()

	q.Push(1, 5)
	q.Push(2, 4)
	q.Push(3, 3)
	q.Push(4, 2)
	q.Push(5, 1)
	q.Push(123, 7)

	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())
	log.Println(q.Shift())

	//q.Pop()
	//q.Pop()
}
