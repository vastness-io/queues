package core

// EmptyQueue represents the value when dequeuing a empty queue.
var EmptyQueue interface{}

// Queue is a abstract data type.
type Queue interface {

	Size() int // Size of the current queue.

	Enqueue(interface{}) // Add node to Tail of the queue.

	Dequeue() interface{} // Removes node from the Head position.
}
