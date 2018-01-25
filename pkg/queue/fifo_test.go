package queue

import (
	"testing"
)

func TestFIFOQueue(t *testing.T) {
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

	if q.Size() != 0 {
		t.Error("Meant to be an empty queue")
	}

}
