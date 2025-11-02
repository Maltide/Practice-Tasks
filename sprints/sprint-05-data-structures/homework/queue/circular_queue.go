package problems

import "fmt"

// CircularQueue represents a circular queue (ring buffer) with fixed capacity
type CircularQueue struct {
	items []int
	head  int
	tail  int
	count int
	size  int
}

// NewCircularQueue creates a new circular queue with the specified capacity
func NewCircularQueue(k int) *CircularQueue {
	// TODO: implement
	return &CircularQueue{size: k}
}

// Push inserts an element into the circular queue
// Returns true if the operation is successful, false if the queue is full
func (cq *CircularQueue) Push(value int) bool {
	// TODO: implement
	if cq.IsFull() {
		return false
	}
	cq.items = append(cq.items, value)
	fmt.Printf("item %v was appended to queue. Queue items now is %v\n", value, cq.items)
	return true
}

// Pull deletes an element from the circular queue
// Returns true if the operation is successful, false if the queue is empty
func (cq *CircularQueue) Pull() bool {
	// TODO: implement
	if cq.IsEmpty() {
		return false
	}
	fmt.Printf("Pull operation: queue now is %v\n", cq.items)
	cq.items = cq.items[1:]
	fmt.Printf("Queue after cutting%v\n", cq.items)
	return true
}

// Front gets the front item from the queue
// Returns the front element or -1 if the queue is empty
func (cq *CircularQueue) Front() int {
	// TODO: implement
	if cq.IsEmpty() {
		return -1
	}
	cq.head = cq.items[0]
	return cq.head
}

// Rear gets the last item from the queue
// Returns the last element or -1 if the queue is empty
func (cq *CircularQueue) Rear() int {
	// TODO: implement
	if cq.IsEmpty() {
		return -1
	}
	cq.tail = cq.items[len(cq.items)-1]
	return cq.tail
}

// IsEmpty checks whether the circular queue is empty
func (cq *CircularQueue) IsEmpty() bool {
	// TODO: implement
	return cq.Size() == 0
}

// IsFull checks whether the circular queue is full
func (cq *CircularQueue) IsFull() bool {
	// TODO: implement
	return cq.size == cq.Size()
}

// Size returns the number of items in the queue
func (cq *CircularQueue) Size() int {
	// TODO: implement
	cq.count = 0
	for range cq.items {
		cq.count++
	}
	return cq.count
}
