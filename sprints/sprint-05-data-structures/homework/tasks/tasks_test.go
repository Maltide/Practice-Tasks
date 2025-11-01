package problems

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// RPN Calculator tests
func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "simple addition",
			tokens:   []string{"2", "1", "+"},
			expected: 3,
		},
		{
			name:     "complex expression",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			name:     "division",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvalRPN(tt.tokens)
			if result != tt.expected {
				t.Errorf("EvalRPN(%v) = %v, want %v", tt.tokens, result, tt.expected)
			}
		})
	}
}

// Valid Parentheses tests
func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "valid simple",
			s:        "()",
			expected: true,
		},
		{
			name:     "valid complex",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "invalid",
			s:        "(]",
			expected: false,
		},
		{
			name:     "valid nested",
			s:        "{[]}",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidParentheses(tt.s)
			if result != tt.expected {
				t.Errorf("IsValidParentheses(%v) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

// StackUsingQueues tests
func TestStackUsingQueues(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "basic operations",
			operations: []string{"push", "push", "top", "pop", "empty"},
			values:     []int{1, 2, 0, 0, 0},
			expected:   []interface{}{nil, nil, 2, 2, false},
		},
		{
			name:       "empty stack operations",
			operations: []string{"empty", "pop"},
			values:     []int{0, 0},
			expected:   []interface{}{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStackUsingQueues()
			for i, op := range tt.operations {
				switch op {
				case "push":
					s.Push(tt.values[i])
				case "pop":
					val, ok := s.Pop()
					if tt.expected[i] == false {
						if ok {
							t.Errorf("Pop() expected failure, got success")
						}
					} else if tt.expected[i] != nil {
						if !ok {
							t.Errorf("Pop() expected success, got failure")
						}
						if val != tt.expected[i] {
							t.Errorf("Pop() = %v, want %v", val, tt.expected[i])
						}
					}
				case "top":
					val, ok := s.Top()
					if tt.expected[i] == false {
						if ok {
							t.Errorf("Top() expected failure, got success")
						}
					} else if tt.expected[i] != nil {
						if !ok {
							t.Errorf("Top() expected success, got failure")
						}
						if val != tt.expected[i] {
							t.Errorf("Top() = %v, want %v", val, tt.expected[i])
						}
					}
				case "empty":
					result := s.Empty()
					if result != tt.expected[i] {
						t.Errorf("Empty() = %v, want %v", result, tt.expected[i])
					}
				}
			}
		})
	}
}

// Sliding Window Maximum tests
func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxSlidingWindow(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

// ReverseList tests
func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "reverse list [1,2,3,4,5]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "reverse list [1,2]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "reverse empty list",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			result := ReverseList(head)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("ReverseList(%v) = %v, want %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

// ReverseListRecursive tests
func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "reverse list [1,2,3,4,5]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "reverse list [1,2]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "reverse empty list",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			result := ReverseListRecursive(head)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("ReverseListRecursive(%v) = %v, want %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

// MergeTwoLists tests
func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "merge [1,2,4] and [1,3,4]",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "merge empty lists",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "merge with one empty",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head1 := createList(tt.list1)
			head2 := createList(tt.list2)
			result := MergeTwoLists(head1, head2)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, resultSlice, tt.expected)
			}
		})
	}
}

// MergeTwoListsRecursive tests
func TestMergeTwoListsRecursive(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "merge [1,2,4] and [1,3,4]",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "merge empty lists",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "merge with one empty",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head1 := createList(tt.list1)
			head2 := createList(tt.list2)
			result := MergeTwoListsRecursive(head1, head2)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("MergeTwoListsRecursive(%v, %v) = %v, want %v", tt.list1, tt.list2, resultSlice, tt.expected)
			}
		})
	}
}

// HasCycle tests
func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int // -1 means no cycle, otherwise index where tail connects
		expected bool
	}{
		{
			name:     "has cycle at position 1",
			values:   []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			name:     "has cycle at position 0",
			values:   []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			name:     "no cycle",
			values:   []int{1},
			pos:      -1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			// Create cycle if pos >= 0
			if tt.pos >= 0 && head != nil {
				tail := head
				cycleNode := head
				for i := 0; i < len(tt.values)-1; i++ {
					if i == tt.pos {
						cycleNode = tail
					}
					tail = tail.Next
				}
				tail.Next = cycleNode
			}
			result := HasCycle(head)
			if result != tt.expected {
				t.Errorf("HasCycle() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// DetectCycle tests
func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int // -1 means no cycle, otherwise index where tail connects
		expected int // expected position of cycle start, -1 if no cycle
	}{
		{
			name:     "cycle at position 1",
			values:   []int{3, 2, 0, -4},
			pos:      1,
			expected: 1,
		},
		{
			name:     "cycle at position 0",
			values:   []int{1, 2},
			pos:      0,
			expected: 0,
		},
		{
			name:     "no cycle",
			values:   []int{1},
			pos:      -1,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			var cycleNode *ListNode
			// Create cycle if pos >= 0
			if tt.pos >= 0 && head != nil {
				tail := head
				current := head
				for i := 0; i < len(tt.values)-1; i++ {
					if i == tt.pos {
						cycleNode = current
					}
					tail = tail.Next
					current = current.Next
				}
				tail.Next = cycleNode
			}
			result := DetectCycle(head)
			if tt.expected == -1 {
				if result != nil {
					t.Errorf("DetectCycle() = %v, want nil", result)
				}
			} else {
				if result != cycleNode {
					t.Errorf("DetectCycle() did not return correct cycle start node")
				}
			}
		})
	}
}
