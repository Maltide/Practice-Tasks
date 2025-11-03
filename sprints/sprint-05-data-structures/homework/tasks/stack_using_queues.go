package problems

import "fmt"

// StackUsingQueues implements a stack using two queues
type StackUsingQueues struct {
	queue1 *Queue
	queue2 *Queue
}

// NewStackUsingQueues creates a new stack implemented using queues
func NewStackUsingQueues() *StackUsingQueues {
	// TODO: implement
	return &StackUsingQueues{queue1: NewQueue(), queue2: NewQueue()}
}

// Push adds an element to the top of the stack
func (s *StackUsingQueues) Push(x int) {
	// TODO: implement
	s.queue2.Push(x)
	fmt.Printf("Queue 2 after start pushing %v\n", s.queue2.items)
	for _, val := range s.queue1.items {
		s.queue2.Push(val)
		fmt.Printf("Copy element of queue1 %v to queue 2. Queue 2 now is %v\n", val, s.queue2.items)
	}
	if s.queue1.Size() != 0 { //стоит ли эту проверку здесь делать чтобы не перезаписывался и так нулевой слайс?
		s.queue1.items = []int{}
	}
	s.queue1, s.queue2 = s.queue2, s.queue1
	fmt.Printf("After pushing queue1 is %v, queue2 is %v\n", s.queue1.items, s.queue2.items)
}

// Pop removes the element on the top of the stack and returns it
// Returns the popped element and a boolean indicating success
func (s *StackUsingQueues) Pop() (int, bool) {
	// TODO: implement
	if s.queue1.IsEmpty() {
		return 0, false
	}
	top, flag := s.queue1.Front()
	s.queue1.items = s.queue1.items[1:]
	return top, flag
}

// Top gets the top element of the stack
// Returns the top element and a boolean indicating success
func (s *StackUsingQueues) Top() (int, bool) {
	// TODO: implement
	if s.queue1.IsEmpty() {
		return 0, false
	}
	return s.queue1.Front()
}

// Empty checks if the stack is empty
func (s *StackUsingQueues) Empty() bool {
	// TODO: implement
	return s.queue1.IsEmpty()
}

// Size returns the number of items in the stack
func (s *StackUsingQueues) Size() int {
	// TODO: implement
	return s.queue1.Size()
}
