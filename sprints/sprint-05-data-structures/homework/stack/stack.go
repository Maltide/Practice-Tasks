package problems

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	// TODO: implement
	return nil
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	// TODO: implement
}

// Pop removes and returns the top item from the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	// TODO: implement
	return 0, nil
}

// Peek returns the top item without removing it
// Returns an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	// TODO: implement
	return 0, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	// TODO: implement
	return false
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	// TODO: implement
	return 0
}
