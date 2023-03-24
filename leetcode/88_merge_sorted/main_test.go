package _88_merge_sorted

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		n     int
		nums2 []int
		m     int
		want  []int
	}{
		{
			name:  "case1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "case1",
			nums1: []int{2, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "case_zero_elems",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, test := range tests {
		l := merge(test.nums1, test.m, test.nums2, test.n)

		require.Equal(t, test.want, l, test.name)
	}
}
