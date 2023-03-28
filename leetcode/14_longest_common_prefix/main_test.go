package _14_longest_common_prefix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name    string
		message string
		got     []string
		want    string
	}{
		{
			name:    "case1",
			message: "the function should return the longest common prefix",
			got:     []string{"flower", "flow", "flight"},
			want:    "fl",
		},
		{
			name:    "case_no_common_prefix",
			message: "if there is no common prefix, an empty string should be returned",
			got:     []string{"dog", "racecar", "car"},
			want:    "",
		},
		{
			name:    "case_single_word_prefix",
			message: "if there is only one word, the world itself should be the longest common prefix",
			got:     []string{"word"},
			want:    "word",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test_result := longestCommonPrefix(test.got)
			require.Equal(t, test.want, test_result, test.message)
		})

	}
}
