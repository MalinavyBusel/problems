package _26

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{
			name: "no duplicates",
			nums: []int{1, 2, 3},
			val:  4,
			want: 3,
		},
		{
			name: "case1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			name: "case2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}

	for _, test := range tests {
		l := removeDuplicate(test.nums, test.val)

		require.Equal(t, test.want, l, test.name)
	}
}
