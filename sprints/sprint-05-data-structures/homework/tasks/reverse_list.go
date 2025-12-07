package problems

// ReverseList reverses a singly linked list
// Returns the new head of the reversed list
func ReverseList(head *ListNode) *ListNode {
	// TODO: implement
	nextNode := (*ListNode)(nil)

	prev := (*ListNode)(nil)

	current := head

	for {
		if current == nil {
			break
		}
		nextNode = current.Next

		current.Next = prev

		prev = current

		current = nextNode
	}
	return prev
}

// ReverseListRecursive reverses a singly linked list using recursion
// Returns the new head of the reversed list
func ReverseListRecursive(head *ListNode) *ListNode {
	// TODO: implement
	if head == nil {
		return nil
	}

	current := head

	if current.Next == nil {
		return current
	}

	newHead := ReverseListRecursive(current.Next)

	current.Next.Next = current

	current.Next = nil

	return newHead
}
