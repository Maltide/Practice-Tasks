package problems

// ListNode represents a node in a singly linked list
// This type is shared across all linked list algorithm tasks
type ListNode struct {
	Val  int
	Next *ListNode
}

// Queue represents a basic FIFO queue
// This type is used by StackUsingQueues
type Queue struct {
	items []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

// Push adds an element to the back of the queue
func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

// Pull removes and returns the element from the front of the queue
func (q *Queue) Pull() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the element at the front without removing it
func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}
