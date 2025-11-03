package problems

import "fmt"

// Queue represents a basic FIFO (First In, First Out) queue
type Queue struct {
	items []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	// TODO: implement
	return &Queue{}
}

// Push adds an element to the back of the queue
func (q *Queue) Push(item int) {
	// TODO: implement
	q.items = append(q.items, item)
	fmt.Printf("Appending item %v. Queue now is %v\n", item, q.items)
}

// Pull removes and returns the element from the front of the queue
// Returns the pulled element and a boolean indicating success
func (q *Queue) Pull() (int, bool) {
	// TODO: implement
	if q.IsEmpty() {
		return 0, false
	}
	fmt.Println("Pulling:")
	front := q.items[0]
	fmt.Printf("front is %v\n", front)
	q.items = q.items[1:]
	fmt.Printf("Cut queue. Now it's%v\n", q.items)
	return front, true
}

// Front returns the element at the front of the queue without removing it
// Returns the front element and a boolean indicating success
func (q *Queue) Front() (int, bool) {
	// TODO: implement
	if q.IsEmpty() {
		return 0, false
	}
	top := q.items[0]
	return top, true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	// TODO: implement
	if len(q.items) == 0 {
		return true
	}
	return false
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	// TODO: implement
	count := 0
	for range q.items {
		count++
	}
	return count
}
