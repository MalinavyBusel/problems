package _39_combination_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		name    string
		message string
		nums    []int
		target  int
		want    [][]int
	}{
		{
			name:    "case1",
			message: "",
			nums:    []int{2, 3, 6, 7},
			target:  7,
			want:    [][]int{{2, 2, 3}, {7}},
		},
		{
			name:    "case2",
			message: "",
			nums:    []int{2, 3, 5},
			target:  8,
			want:    [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:    "case_no_combinations",
			message: "if there is no valid combinations, return an empty list",
			nums:    []int{2},
			target:  1,
			want:    [][]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testResult := combinationSum(test.nums, test.target)
			if !assert.ElementsMatch(t, test.want, testResult) {
				t.Errorf("\n got: %v,\n want: %v,\n message: %s", testResult, test.want, test.message)
			}
		})

	}
}
