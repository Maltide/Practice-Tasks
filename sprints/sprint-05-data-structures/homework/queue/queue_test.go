package problems

import (
	"testing"
)

// Queue tests
func TestQueue(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "basic operations",
			operations: []string{"push", "push", "pull", "push", "front", "size"},
			values:     []int{1, 2, 0, 3, 0, 0},
			expected:   []interface{}{nil, nil, 1, nil, 2, 2},
		},
		{
			name:       "empty queue operations",
			operations: []string{"isempty", "size", "pull", "front"},
			values:     []int{0, 0, 0, 0},
			expected:   []interface{}{true, 0, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for i, op := range tt.operations {
				switch op {
				case "push":
					q.Push(tt.values[i])
				case "pull":
					val, ok := q.Pull()
					if tt.expected[i] == false {
						if ok {
							t.Errorf("Pull() expected failure, got success")
						}
					} else if tt.expected[i] != nil {
						if !ok {
							t.Errorf("Pull() expected success, got failure")
						}
						if val != tt.expected[i] {
							t.Errorf("Pull() = %v, want %v", val, tt.expected[i])
						}
					}
				case "front":
					val, ok := q.Front()
					if tt.expected[i] == false {
						if ok {
							t.Errorf("Front() expected failure, got success")
						}
					} else if tt.expected[i] != nil {
						if !ok {
							t.Errorf("Front() expected success, got failure")
						}
						if val != tt.expected[i] {
							t.Errorf("Front() = %v, want %v", val, tt.expected[i])
						}
					}
				case "isempty":
					result := q.IsEmpty()
					if result != tt.expected[i] {
						t.Errorf("IsEmpty() = %v, want %v", result, tt.expected[i])
					}
				case "size":
					result := q.Size()
					if result != tt.expected[i] {
						t.Errorf("Size() = %v, want %v", result, tt.expected[i])
					}
				}
			}
		})
	}
}

// CircularQueue tests
func TestCircularQueue(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "basic operations",
			capacity:   3,
			operations: []string{"push", "push", "push", "push", "pull", "push", "front", "rear"},
			values:     []int{1, 2, 3, 4, 0, 4, 0, 0},
			expected:   []interface{}{true, true, true, false, true, true, 2, 4},
		},
		{
			name:       "empty queue operations",
			capacity:   2,
			operations: []string{"isempty", "pull", "front"},
			values:     []int{0, 0, 0},
			expected:   []interface{}{true, false, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := NewCircularQueue(tt.capacity)
			for i, op := range tt.operations {
				switch op {
				case "push":
					result := cq.Push(tt.values[i])
					if result != tt.expected[i] {
						t.Errorf("Push(%v) = %v, want %v", tt.values[i], result, tt.expected[i])
					}
				case "pull":
					result := cq.Pull()
					if result != tt.expected[i] {
						t.Errorf("Pull() = %v, want %v", result, tt.expected[i])
					}
				case "front":
					result := cq.Front()
					if result != tt.expected[i] {
						t.Errorf("Front() = %v, want %v", result, tt.expected[i])
					}
				case "rear":
					result := cq.Rear()
					if result != tt.expected[i] {
						t.Errorf("Rear() = %v, want %v", result, tt.expected[i])
					}
				case "isempty":
					result := cq.IsEmpty()
					if result != tt.expected[i] {
						t.Errorf("IsEmpty() = %v, want %v", result, tt.expected[i])
					}
				case "isfull":
					result := cq.IsFull()
					if result != tt.expected[i] {
						t.Errorf("IsFull() = %v, want %v", result, tt.expected[i])
					}
				}
			}
		})
	}
}
