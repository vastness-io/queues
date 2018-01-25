package queue

import (
	"sync"
	"github.com/vastness-io/queues/pkg/core"
)

// FIFOQueue is a first-In-First-Out implementation of Queue.
type FIFOQueue struct {
	head  int // head is the first index of the slice.
	nodes []interface{}
	cond  *sync.Cond
	count int
	tail  int // tail is the last index of the slice.
}

// NewFIFOQueue creates a new First-In-First-Out Queue.
func NewFIFOQueue() core.Queue {
	return &FIFOQueue{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// Size of the queue.
func (q *FIFOQueue) Size() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.count
}

// Enqueue adds the node to the queue.
func (q *FIFOQueue) Enqueue(node interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.nodes = append(q.nodes, node)
	q.count++
	q.tail++
	q.cond.Signal()
}

// Dequeue removes and returns the node from the Head of the queue.
func (q *FIFOQueue) Dequeue() interface{} {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if q.count == 0 {
		q.cond.Wait()
	}

	head := q.nodes[q.head]
	q.nodes = q.nodes[1:]
	q.count--
	q.tail--
	return head
}
