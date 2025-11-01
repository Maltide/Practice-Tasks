package problems

import (
	"testing"
)

// Stack tests
func TestStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "basic operations",
			operations: []string{"push", "push", "pop", "push", "peek", "size"},
			values:     []int{1, 2, 0, 3, 0, 0},
			expected:   []interface{}{nil, nil, 2, nil, 3, 2},
		},
		{
			name:       "empty stack operations",
			operations: []string{"isempty", "size", "pop", "peek"},
			values:     []int{0, 0, 0, 0},
			expected:   []interface{}{true, 0, "error", "error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for i, op := range tt.operations {
				switch op {
				case "push":
					s.Push(tt.values[i])
				case "pop":
					val, err := s.Pop()
					if tt.expected[i] == "error" {
						if err == nil {
							t.Errorf("Pop() expected error, got nil")
						}
					} else if tt.expected[i] != nil {
						if err != nil {
							t.Errorf("Pop() error = %v, expected no error", err)
						}
						if val != tt.expected[i] {
							t.Errorf("Pop() = %v, want %v", val, tt.expected[i])
						}
					}
				case "peek":
					val, err := s.Peek()
					if tt.expected[i] == "error" {
						if err == nil {
							t.Errorf("Peek() expected error, got nil")
						}
					} else if tt.expected[i] != nil {
						if err != nil {
							t.Errorf("Peek() error = %v, expected no error", err)
						}
						if val != tt.expected[i] {
							t.Errorf("Peek() = %v, want %v", val, tt.expected[i])
						}
					}
				case "isempty":
					result := s.IsEmpty()
					if result != tt.expected[i] {
						t.Errorf("IsEmpty() = %v, want %v", result, tt.expected[i])
					}
				case "size":
					result := s.Size()
					if result != tt.expected[i] {
						t.Errorf("Size() = %v, want %v", result, tt.expected[i])
					}
				}
			}
		})
	}
}

// MinStack tests
func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "basic min operations",
			operations: []string{"push", "push", "push", "getmin", "pop", "getmin"},
			values:     []int{3, 1, 2, 0, 0, 0},
			expected:   []interface{}{nil, nil, nil, 1, 2, 1},
		},
		{
			name:       "empty stack min",
			operations: []string{"getmin"},
			values:     []int{0},
			expected:   []interface{}{"error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMinStack()
			for i, op := range tt.operations {
				switch op {
				case "push":
					ms.Push(tt.values[i])
				case "pop":
					val, ok := ms.Pop()
					if tt.expected[i] == "error" {
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
				case "getmin":
					val, ok := ms.GetMin()
					if tt.expected[i] == "error" {
						if ok {
							t.Errorf("GetMin() expected failure, got success")
						}
					} else if tt.expected[i] != nil {
						if !ok {
							t.Errorf("GetMin() expected success, got failure")
						}
						if val != tt.expected[i] {
							t.Errorf("GetMin() = %v, want %v", val, tt.expected[i])
						}
					}
				}
			}
		})
	}
}
