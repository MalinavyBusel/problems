package _46_permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name    string
		message string
		nums    []int
		want    [][]int
	}{
		{
			name:    "case1",
			message: `should be returned "len(nums)*(len(nums)-1)*...*1 = 6 unique permutations in any order"`,
			nums:    []int{1, 2, 3},
			want:    [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:    "case2",
			message: `should be returned "len(nums)*(len(nums)-1) = 2*1 = 1 unique permutations in any order"`,
			nums:    []int{0, 1},
			want:    [][]int{{0, 1}, {1, 0}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testResult := permute(test.nums)
			if !assert.ElementsMatch(t, test.want, testResult) {
				t.Errorf("\n got: %v,\n want: %v,\n message: %s", testResult, test.want, test.message)
			}
		})

	}
}
