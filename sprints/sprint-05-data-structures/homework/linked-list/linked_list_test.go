package problems

import (
	"testing"
)

// ListNode tests
func TestListNode(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		expected int
	}{
		{
			name:     "create node with value 5",
			val:      5,
			expected: 5,
		},
		{
			name:     "create node with value 0",
			val:      0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewListNode(tt.val)
			if node == nil {
				t.Error("NewListNode() returned nil")
				return
			}
			if node.Val != tt.expected {
				t.Errorf("NewListNode(%v).Val = %v, want %v", tt.val, node.Val, tt.expected)
			}
		})
	}
}

// LinkedList tests
func TestLinkedList(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		indices    []int
		expected   []interface{}
	}{
		{
			name:       "basic operations",
			operations: []string{"addathead", "addattail", "addatindex", "get", "deleteatindex", "get"},
			values:     []int{1, 3, 1, 2, 0, 1, 0},
			indices:    []int{0, 0, 1, 1, 1, 1},
			expected:   []interface{}{nil, nil, nil, 2, true, 3},
		},
		{
			name:       "empty list operations",
			operations: []string{"isempty", "size", "get"},
			values:     []int{0, 0, 0},
			indices:    []int{0, 0, 0},
			expected:   []interface{}{true, 0, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()
			for i, op := range tt.operations {
				switch op {
				case "addathead":
					ll.AddAtHead(tt.values[i])
				case "addattail":
					ll.AddAtTail(tt.values[i])
				case "addatindex":
					ll.AddAtIndex(tt.indices[i], tt.values[i])
				case "deleteatindex":
					result := ll.DeleteAtIndex(tt.indices[i])
					if result != tt.expected[i] {
						t.Errorf("DeleteAtIndex(%v) = %v, want %v", tt.indices[i], result, tt.expected[i])
					}
				case "get":
					result := ll.Get(tt.indices[i])
					if result != tt.expected[i] {
						t.Errorf("Get(%v) = %v, want %v", tt.indices[i], result, tt.expected[i])
					}
				case "isempty":
					result := ll.IsEmpty()
					if result != tt.expected[i] {
						t.Errorf("IsEmpty() = %v, want %v", result, tt.expected[i])
					}
				case "size":
					result := ll.Size()
					if result != tt.expected[i] {
						t.Errorf("Size() = %v, want %v", result, tt.expected[i])
					}
				}
			}
		})
	}
}
