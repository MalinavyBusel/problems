package _20_valid_parentheses

func isValid(s string) bool {
	var stack = make([]byte, 0, 5000)

	bracketPair := map[byte]byte{'}': '{', ']': '[', ')': '('}

	for _, x := range s {
		switch x {
		case '[', '{', '(':
			{
				stack = append(stack, byte(x))
			}
		default:
			{
				if len(stack) == 0 {
					return false
				}

				if stack[len(stack)-1] != bracketPair[byte(x)] {
					return false
				}

				stack = stack[:len(stack)-1]

			}
		}
	}
	return len(stack) == 0
}
