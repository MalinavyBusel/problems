package _0_remove_duplicates_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "no duplicates",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "case1",
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
		},
		{
			name: "case2",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
	}

	for _, test := range tests {
		l := removeDuplicates(test.nums)

		require.Equal(t, test.want, l, test.name)
	}
}
