package _383_ransom_note

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		ransom   string
		magazine string
		want     bool
	}{
		{
			name:     "case1",
			message:  "there is not enough letters b in the magazine string",
			ransom:   "aa",
			magazine: "ab",
			want:     false,
		},
		{
			name: "case2",
			message: "ransom string can be constructed from magazine, because it contains at least the needed amount of " +
				"letters from ransom",
			ransom:   "aa",
			magazine: "aab",
			want:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test_result := canConstruct(test.ransom, test.magazine)
			require.Equal(t, test.want, test_result, test.message)
		})

	}
}
