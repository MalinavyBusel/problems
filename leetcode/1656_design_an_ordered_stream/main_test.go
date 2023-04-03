package _1656_design_an_ordered_stream

import (
	"testing"
)

func Test_OrderedStreamInsert(t *testing.T) {
	stream := Constructor(5)
	returned := stream.Insert(3, "ccccc") // Inserts (3, "ccccc"), returns [].
	if len(returned) != 0 {
		t.Errorf("got len: %q, want: %q", len(returned), 0)
	}
	returned = stream.Insert(1, "aaaaa") // Inserts (1, "aaaaa"), returns ["aaaaa"].
	if len(returned) != 1 {
		t.Errorf("got len: %q, want: %q", len(returned), 1)
	}
	returned = stream.Insert(2, "bbbbb") // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
	if len(returned) != 2 {
		t.Errorf("got len: %q, want: %q", len(returned), 2)
	}
	returned = stream.Insert(5, "eeeee") // Inserts (5, "eeeee"), returns [].
	if len(returned) != 0 {
		t.Errorf("got len: %q, want: %q", len(returned), 0)
	}
	returned = stream.Insert(4, "ddddd") // Inserts (4, "ddddd"), returns ["ddddd", "eeeee"].
	if len(returned) != 2 {
		t.Errorf("got len: %q, want: %q", len(returned), 2)
	}
}
