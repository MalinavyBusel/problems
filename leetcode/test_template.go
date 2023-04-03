package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name    string
		message string
		got     string
		want    bool
	}{
		{
			name:    "case1",
			message: "every close bracket has a corresponding open bracket of the same type",
			got:     "()",
			want:    true,
		},
		{
			name:    "case2",
			message: "many types of parentheses should be accepted",
			got:     "()[]{}",
			want:    true,
		},
		{
			name:    "case_different_opening_and_closing_parentheses",
			message: "open brackets must be closed by the same type of brackets.",
			got:     "(]",
			want:    false,
		},
		{
			name:    "case_wrong_order_of_closing",
			message: "open brackets must be closed in the correct order",
			got:     "({)}",
			want:    false,
		},
		{
			name:    "case_not_all_parentheses_closed",
			message: "every opened parenthesis should be closed",
			got:     "((()",
			want:    false,
		},
		{
			name:    "case_closed_without_opening",
			message: "you cannot close the bracket without opening it",
			got:     ")()",
			want:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test_result := string((test.got))
			require.Equal(t, test.want, test_result, test.message)
		})

	}
}
