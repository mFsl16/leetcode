package mergeshortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	output := ListNode{}
	output.Next = nil
	tail := &output

	for list1 != nil || list2 != nil {
		if list1 == nil {
			(*tail).Next = list2
			break
		} else if list2 == nil {
			(*tail).Next = list1
			break
		} else if list1.Val <= list2.Val {
			(*tail).Next = list1
			list1 = (*list1).Next
		} else {
			(*tail).Next = list2
			list2 = (*list2).Next
		}
		tail = tail.Next
	}

	return output.Next
}
