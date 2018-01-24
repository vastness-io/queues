package queue

import (
	"sync"
)

//EmptyQueue represents the value when dequeuing a empty queue
var EmptyQueue interface {}

//Queue is an sequentially abstract data type
type Queue interface {
	//Size of the current queue
	Size() int
	//Add node to Tail of the queue
	Enqueue(interface{})
	//Removes node from the Head position
	Dequeue() interface{}
}

//FIFOQueue is a first-In-First-Out implementation of Queue
type FIFOQueue struct {
	//head is the first index of the slice
	head int
	nodes []interface{}
	sync.RWMutex
	count int
	//tail is the last index of the slice
	tail int
}

//NewFIFOQueue creates a new First-In-First-Out Queue
func NewFIFOQueue() Queue {
	return &FIFOQueue{}
}

//Size of the queue
func (q *FIFOQueue) Size() int {
	q.RLock()
	defer q.RUnlock()
	return q.count
}

//Enqueue adds the node to the queue
func (q *FIFOQueue) Enqueue(node interface{}) {
	q.Lock()
	q.nodes = append(q.nodes, node)
	q.count++
	q.tail++
	q.Unlock()
}

//Dequeue removes and returns the node from the Head of the queue
func (q *FIFOQueue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()
	
	if q.count > 0 {
		head := q.nodes[q.head]
		q.nodes = q.nodes[q.head+1:]
		q.count--
		q.tail--
		return head
	}

	return EmptyQueue
}
