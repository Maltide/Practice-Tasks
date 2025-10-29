package problems

// MinStack represents a stack that supports retrieving the minimum element in O(1) time
type MinStack struct {
	stack    []int // Main stack to store all elements
	minStack []int // Auxiliary stack to store minimum elements
}

// NewMinStack creates a new MinStack
func NewMinStack() *MinStack {
	// TODO: implement
	return nil
}

// Push adds an element to the top of the stack
func (ms *MinStack) Push(val int) {
	// TODO: implement
}

// Pop removes the element on the top of the stack
// Returns the popped value and a boolean indicating success
func (ms *MinStack) Pop() (int, bool) {
	// TODO: implement
	return 0, false
}

// Top gets the top element of the stack
// Returns the top element and a boolean indicating success
func (ms *MinStack) Top() (int, bool) {
	// TODO: implement
	return 0, false
}

// GetMin retrieves the minimum element in the stack
// Returns the minimum element and a boolean indicating success
func (ms *MinStack) GetMin() (int, bool) {
	// TODO: implement
	return 0, false
}

// IsEmpty checks if the stack is empty
func (ms *MinStack) IsEmpty() bool {
	// TODO: implement
	return false
}

// Size returns the number of items in the stack
func (ms *MinStack) Size() int {
	// TODO: implement
	return 0
}
