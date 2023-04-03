package _203_remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	first := &ListNode{}
	last := first
	for node := head; node != nil; node = node.Next {
		if node.Val != val {
			last.Next = node
			last = node
		}
	}
	last.Next = nil
	return first.Next
}
