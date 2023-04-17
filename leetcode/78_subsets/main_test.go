package _78_subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		name    string
		message string
		nums    []int
		want    [][]int
	}{
		{
			name:    "case1",
			message: `return all the subsets of the given set, the solution may be returned in any order`,
			nums:    []int{1, 2, 3},
			want:    [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:    "case2",
			message: "zero subset also should be included",
			nums:    []int{0},
			want:    [][]int{{}, {0}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testResult := subsets(test.nums)
			if !assert.ElementsMatch(t, test.want, testResult) {
				t.Errorf("\n got: %v,\n want: %v,\n message: %s", testResult, test.want, test.message)
			}
		})

	}
}
