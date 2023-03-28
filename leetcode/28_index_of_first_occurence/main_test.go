package _28_index_of_first_occurence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "case1",
			message:  "The first occurrence is at index 0, so 0 should be returned",
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			name:     "case2",
			message:  "'leeto' did not occur in 'leetcode', so we should return -1.",
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test_result := strStr(test.haystack, test.needle)
			require.Equal(t, test.want, test_result, test.message)
		})

	}
}
