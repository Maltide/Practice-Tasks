package problems

// MergeTwoLists merges two sorted linked lists into one sorted list
// Returns the head of the merged linked list
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: implement
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	resList := &ListNode{}

	current := resList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1

			list1 = list1.Next

			current = current.Next
		} else {
			current.Next = list2

			list2 = list2.Next

			current = current.Next
		}
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return resList.Next
}

// MergeTwoListsRecursive merges two sorted linked lists recursively
// Returns the head of the merged linked list
func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: implement
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list2.Next, list1)
		return list2
	}
}
