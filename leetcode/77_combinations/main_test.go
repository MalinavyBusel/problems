package _77_combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		name    string
		message string
		n       int
		k       int
		want    [][]int
	}{
		{
			name: "case1",
			message: `there are 4 choose 2 = 6 total combinations. 
Note that combinations are unordered, i.e.,
[1,2] and [2,1] are considered to be the same combination.`,
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:    "case2",
			message: "there is 1 choose 1 = 1 total combination.",
			n:       1,
			k:       1,
			want:    [][]int{{1}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testResult := combine(test.n, test.k)
			if !reflect.DeepEqual(test.want, testResult) {
				t.Errorf("\n got: %v,\n want: %v,\n message: %s", testResult, test.want, test.message)
			}
		})

	}
}
