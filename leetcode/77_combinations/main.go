package _77_combinations

func combine(n int, k int) [][]int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i + 1
	}

	result := [][]int{}
	buf := make([]int, 0, k)

	var backtrack func(int, []int, []int)
	backtrack = func(left int, current []int, numbers []int) {
		if left <= 0 {
			newComb := append([]int{}, current...)
			result = append(result, newComb)
			return
		}
		for i, _ := range numbers {
			backtrack(left-1, append(current, numbers[i]), numbers[i+1:])
		}
	}
	backtrack(k, buf, numbers)
	return result
}
