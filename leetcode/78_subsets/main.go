package _78_subsets

func subsets(nums []int) [][]int {

	result := [][]int{}
	buf := make([]int, 0, len(nums))

	var backtrack func(int, []int, []int)
	backtrack = func(left int, current []int, numbers []int) {
		newComb := append([]int{}, current...)
		result = append(result, newComb)
		if left <= 0 {
			return
		}
		for i, _ := range numbers {
			backtrack(left-1, append(current, numbers[i]), numbers[i+1:])
		}
	}
	backtrack(len(nums), buf, nums)
	return result
}
