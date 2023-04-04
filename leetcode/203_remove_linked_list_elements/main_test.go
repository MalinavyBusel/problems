package _203_remove_linked_list_elements

import "testing"

func Test_MultipleNodes(t *testing.T) {
	seven := ListNode{Val: 6}
	six := ListNode{Val: 5, Next: &seven}
	five := ListNode{Val: 4, Next: &six}
	four := ListNode{Val: 3, Next: &five}
	three := ListNode{Val: 6, Next: &four}
	two := ListNode{Val: 2, Next: &three}
	one := ListNode{Val: 1, Next: &two}
	res := removeElements(&one, 1)
	five_ := ListNode{Val: 5}
	four_ := ListNode{Val: 4, Next: &five_}
	three_ := ListNode{Val: 3, Next: &four_}
	two_ := ListNode{Val: 2, Next: &three_}
	one_ := ListNode{Val: 1, Next: &two_}
	for (node_1 := one; node_2 := res; i := 0); i<5; i++ {
		if node.Val != val {
			last.Next = node
			last = node
		}
	}
	if *res != want {
		t.Error(res, " ", want)
	}
}
