# Linked List Implementation Tasks

This directory contains linked list problems and implementations.

## Files

### 1. linked_list.go
Basic singly linked list implementation with the following operations:
- `Insert(val int, pos int) error` - Insert value at position
- `Delete(pos int) error` - Delete node at position
- `Find(val int) int` - Find first occurrence of value
- `Size() int` - Get number of nodes
- `ToSlice() []int` - Convert list to slice (for testing)

Node structure:
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

### 2. reverse_list.go
**LeetCode 206: Reverse Linked List**

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example:
- Input: head = [1,2,3,4,5]
- Output: [5,4,3,2,1]

Can be solved iteratively or recursively.

### 3. detect_cycle.go
**LeetCode 141: Linked List Cycle**

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.

Use Floyd's Cycle-Finding Algorithm (Tortoise and Hare).

### 4. merge_sorted_lists.go
**LeetCode 21: Merge Two Sorted Lists**

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

## Implementation Notes

- Handle edge cases: empty lists, single node, invalid positions
- Use proper pointer manipulation
- Avoid memory leaks
- Follow Go conventions for error handling
- For cycle detection, achieve O(1) space complexity

## Common Patterns

### Creating a Linked List from Array
```go
func arrayToList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{Val: arr[0]}
    current := head
    for i := 1; i < len(arr); i++ {
        current.Next = &ListNode{Val: arr[i]}
        current = current.Next
    }
    return head
}
```

### Traversing a Linked List
```go
current := head
for current != nil {
    // Process current.Val
    current = current.Next
}
```

## Testing

Run tests with:
```bash
go test ./sprints/sprint-05-data-structures/homework -run TestLinkedList