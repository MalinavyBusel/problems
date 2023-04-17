package _39_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	current := []int{}

	var backtrack func(int, []int)
	backtrack = func(start int, current []int) {
		if sum := arrSum(current); sum == target {
			newArr := append([]int{}, current...)
			res = append(res, newArr)
		} else if sum > target {
			return
		}

		for i, _ := range candidates[start:] {
			backtrack(start+i, append(current, candidates[start+i]))
		}
	}
	backtrack(0, current)
	return res
}

func arrSum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}
