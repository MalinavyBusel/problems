package _21_merge_sorted_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{}
	last := &head
	for list1 != nil {
		if list2 != nil && list2.Val < list1.Val {
			last.Next = list2
			last = list2
			list2 = list2.Next
		} else {
			last.Next = list1
			last = list1
			list1 = list1.Next
		}
	}
	for list2 != nil {
		last.Next = list2
		last = list2
		list2 = list2.Next
	}
	return head.Next
}
