package problems

// MinStack represents a stack that supports retrieving the minimum element in O(1) time
type MinStack struct {
	stack    []int // Main stack to store all elements
	minStack []int // Auxiliary stack to store minimum elements
}

// NewMinStack creates a new MinStack
func NewMinStack() *MinStack {
	// TODO: implement
	return &MinStack{stack: []int{}, minStack: []int{}}
}

// Push adds an element to the top of the stack
func (ms *MinStack) Push(val int) {
	// TODO: implement
	ms.stack = append(ms.stack, val)
	top, flag := ms.Top()
	if !flag {
		ms.minStack = append(ms.minStack, val)
	} else {
		if val <= top {
			ms.minStack = append(ms.minStack, val)
		} else {
			ms.minStack = append(ms.minStack, top)
		}
	}
}

// Pop removes the element on the top of the stack
// Returns the popped value and a boolean indicating success
func (ms *MinStack) Pop() (int, bool) {
	// TODO: implement
	if ms.IsEmpty() {
		return 0, false
	}
	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
	return top, true
}

// Top gets the top element of the stack
// Returns the top element and a boolean indicating success
func (ms *MinStack) Top() (int, bool) {
	// TODO: implement
	if ms.IsEmpty() {
		return 0, false
	}
	top := ms.minStack[len(ms.minStack)-1]
	return top, true
}

// GetMin retrieves the minimum element in the stack
// Returns the minimum element and a boolean indicating success
func (ms *MinStack) GetMin() (int, bool) {
	// TODO: implement
	if ms.IsEmpty() {
		return 0, false
	}
	minelem, _ := ms.Top()
	return minelem, true
}

// IsEmpty checks if the stack is empty
func (ms *MinStack) IsEmpty() bool {
	// TODO: implement
	if len(ms.stack) == 0 || len(ms.minStack) == 0 {
		return true
	}
	return false
}

// Size returns the number of items in the stack
func (ms *MinStack) Size() int {
	// TODO: implement
	count := 0
	for range ms.stack {
		count++
	}
	return count
}
