package _28_index_of_first_occurence

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}

	return -1
}
