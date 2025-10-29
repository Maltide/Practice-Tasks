package problems

// StackUsingQueues implements a stack using two queues
type StackUsingQueues struct {
	queue1 *Queue
	queue2 *Queue
}

// NewStackUsingQueues creates a new stack implemented using queues
func NewStackUsingQueues() *StackUsingQueues {
	// TODO: implement
	return nil
}

// Push adds an element to the top of the stack
func (s *StackUsingQueues) Push(x int) {
	// TODO: implement
}

// Pop removes the element on the top of the stack and returns it
// Returns the popped element and a boolean indicating success
func (s *StackUsingQueues) Pop() (int, bool) {
	// TODO: implement
	return 0, false
}

// Top gets the top element of the stack
// Returns the top element and a boolean indicating success
func (s *StackUsingQueues) Top() (int, bool) {
	// TODO: implement
	return 0, false
}

// Empty checks if the stack is empty
func (s *StackUsingQueues) Empty() bool {
	// TODO: implement
	return false
}

// Size returns the number of items in the stack
func (s *StackUsingQueues) Size() int {
	// TODO: implement
	return 0
}
