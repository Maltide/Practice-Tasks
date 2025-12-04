package problems

import "fmt"

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *ListNode
	size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	// TODO: implement
	newLink := &LinkedList{}
	return newLink
}

// NewListNode creates a new list node with the given value
func NewListNode(val int) *ListNode {
	// TODO: implement
	listNode := ListNode{Val: val}
	return &listNode
}

// AddAtHead adds a node of value val before the first element
func (ll *LinkedList) AddAtHead(val int) {
	// TODO: implement
	newnode := NewListNode(val)
	newnode.Next = ll.head
	ll.head = newnode
	ll.size++
}

// AddAtTail appends a node of value val as the last element
func (ll *LinkedList) AddAtTail(val int) {
	// TODO: implement
	if ll.IsEmpty() {
		newNode := NewListNode(val)
		ll.head = newNode
		ll.size++
		return
	}

	current := ll.head

	newNode := NewListNode(val)

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.size++
}

// AddAtIndex adds a node of value val before the index-th node
// If index equals the length of the list, the node is appended to the end
// If index is greater than the length, the node is not inserted
func (ll *LinkedList) AddAtIndex(index int, val int) {
	// TODO: implement
	if index == ll.size {
		ll.AddAtTail(val)
		return
	}

	if index == 0 {
		ll.AddAtHead(val)
		return
	}

	if index > ll.size || index < 0 {
		fmt.Println("Index is out of range LinkedList")
		return
	}

	newNode := NewListNode(val)

	preNode := (*ListNode)(nil)

	counter := 0

	currentNode := ll.head

	for counter < index {

		preNode = currentNode

		currentNode = currentNode.Next

		counter++

	}
	newNode.Next = currentNode

	preNode.Next = newNode

	ll.size++
}

// DeleteAtIndex deletes the index-th node in the linked list
// Returns true if the deletion was successful, false otherwise
func (ll *LinkedList) DeleteAtIndex(index int) bool {
	// TODO: implement
	if index >= ll.size || index < 0 {
		return false
	}

	if index == 0 {
		ll.head = ll.head.Next
		ll.size--
		return true
	}

	counter := 0

	preNode := (*ListNode)(nil)

	currentNode := ll.head

	for {
		if counter == index {
			preNode.Next = currentNode.Next

			ll.size--

			break
		}

		preNode = currentNode

		currentNode = currentNode.Next

		counter++
	}

	return true
}

// Get returns the value of the index-th node in the linked list
// Returns -1 if the index is invalid
func (ll *LinkedList) Get(index int) int {
	// TODO: implement
	if index < 0 || index >= ll.size {
		return -1
	}

	counter := 0

	current := ll.head

	for {
		if counter == index {
			return current.Val
		}
		current = current.Next
		counter++
	}
}

// IsEmpty checks if the linked list is empty
func (ll *LinkedList) IsEmpty() bool {
	// TODO: implement
	return ll.size == 0
}

// Size returns the number of elements in the linked list
func (ll *LinkedList) Size() int {
	// TODO: implement
	return ll.size
}

// ToSlice converts the linked list to a slice of integers
func (ll *LinkedList) ToSlice() []int {
	// TODO: implement
	if ll.IsEmpty() {
		return nil
	}

	current := ll.head

	rslice := make([]int, 0, ll.size)

	for current != nil {
		rslice = append(rslice, current.Val)

		current = current.Next
	}
	return rslice
}
