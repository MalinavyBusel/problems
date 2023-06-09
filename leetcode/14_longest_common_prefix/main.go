package _14_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		letter := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != letter {
				return strs[j][:i]
			}
		}
	}
	return strs[0]
}
