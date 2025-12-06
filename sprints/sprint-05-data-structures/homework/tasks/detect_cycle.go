package problems

// HasCycle checks if a linked list has a cycle
// Uses Floyd's Cycle Detection Algorithm (Tortoise and Hare)
func HasCycle(head *ListNode) bool {
	// TODO: implement
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head

	fast := head

	for {
		slow = slow.Next

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}
}

// DetectCycle finds the node where the cycle in a linked list begins
// Returns nil if there is no cycle
// Uses Floyd's Cycle Detection Algorithm with additional step to find cycle start
func DetectCycle(head *ListNode) *ListNode {
	// TODO: implement
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head

	fast := head

	for {
		slow = slow.Next

		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		fast = fast.Next.Next

		if slow == fast {
			ptr1 := head
			ptr2 := slow

			for ptr1 != ptr2 {
				ptr1 = ptr1.Next
				ptr2 = ptr2.Next
			}
			return ptr1
		}
	}
}
