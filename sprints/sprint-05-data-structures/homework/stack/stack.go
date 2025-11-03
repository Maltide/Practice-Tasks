// package problems contains basic stack-functions and tests for them
package problems

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	// TODO: implement
	return &Stack{items: []int{}}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	// TODO: implement
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	// TODO: implement
	top, err := s.Peek()
	if err != nil {
		return 0, err
	}
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek returns the top item without removing it
// Returns an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	// TODO: implement
	if s.IsEmpty() {
		return 0, fmt.Errorf("empty stack")
	}
	top := s.items[len(s.items)-1]
	return top, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	// TODO: implement
	if len(s.items) == 0 {
		return true
	}
	return false
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	// TODO: implement
	count := 0
	for range s.items {
		count++
	}
	return count
}
