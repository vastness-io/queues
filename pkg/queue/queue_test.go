package queue

import (
	"testing"
)
func TestEnqueue(t *testing.T) {

	fifo := FIFOQueue{}

	if fifo.Size() != 0 {
		t.Fail()
	}

	fifo.Enqueue(1)
	if fifo.Size() != 1 && fifo.tail != 1{
		t.Fail()
	}
	
	fifo.Enqueue(2)

	ele := fifo.Dequeue() 

	v := ele.(int)

	if v != 1 {
		t.Fail()
	}

}

func TestQueue(t *testing.T) {
	q := NewFIFOQueue()
	count := 10000
	var first, tail int
	for i := 0; i < count; i++ {
		q.Enqueue(i)
		tail++
		if q.Size() <= 0 {
			t.Error("Shouldn't be empty")
		}
		if q.Size() != tail {
			t.Error("Should be equal")
		}
		if q.Dequeue().(int) != first {
			t.Error("Should be the same element")
		}

		tail--
		first++
	}

	if q.Dequeue() != EmptyQueue {
		t.Error("Meant to be an empty queue")
	}
	
}