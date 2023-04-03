package _1290_binary_to_int

import "testing"

func Test_SingleNode(t *testing.T) {
	want := 1
	c := ListNode{Val: 1}
	res := getDecimalValue(&c)
	if res != want {
		t.Errorf("got: %q, want: %q", res, want)
	}
}

func Test_MultipleNodes(t *testing.T) {
	want := 5
	c := ListNode{Val: 1}
	b := ListNode{
		Val:  0,
		Next: &c,
	}
	a := ListNode{Val: 1, Next: &b}
	res := getDecimalValue(&a)
	if res != want {
		t.Errorf("got: %q, want: %q", res, want)
	}
}
