package _26

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want int
	}{
		{
			name: "no duplicates",
			got:  []int{1, 2, 3},
			want: 3,
		},
		{
			name: "case1",
			got:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		{
			name: "case2",
			got:  []int{1, 1, 2},
			want: 2,
		},
	}

	for _, test := range tests {
		l := removeDuplicates(test.got)

		require.Equal(t, test.want, l)
	}
}
